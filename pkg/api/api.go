package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/istyf/siusmitm/pkg/application"
	"github.com/istyf/siusmitm/pkg/components"
	"github.com/istyf/siusmitm/pkg/mitm"

	. "github.com/diwise/frontend-toolkit"
	"github.com/diwise/frontend-toolkit/pkg/assets"
	"github.com/diwise/service-chassis/pkg/infrastructure/net/http/router"
	"github.com/diwise/service-chassis/pkg/infrastructure/o11y/logging"
	"github.com/google/uuid"
)

func NewAssetLoader(ctx context.Context, assetPath string) (AssetLoader, error) {
	logging.GetFromContext(ctx).Info("creating asset loader", "path", assetPath)

	return assets.NewLoader(ctx,
		assets.BasePath(assetPath), assets.Exclude("/l10n"),
		assets.Logger(logging.GetFromContext(ctx)),
	)
}

func RegisterHandlers(ctx context.Context, mux2 *http.ServeMux, middleware []func(http.Handler) http.Handler, assetLoader AssetLoader, app *application.App) error {

	version := uuid.NewString()

	r := router.New(mux2)

	assets.RegisterEndpoints(ctx, assetLoader, assets.WithMux(mux2),
		assets.WithImmutableExpiry(48*time.Hour),
		assets.WithRedirect("/favicon.ico", "/icons/favicon.ico", http.StatusFound),
	)

	shots := make([]mitm.Shot, 0, 100)
	//json.Unmarshal([]byte(shotsJSON), &shots)

	mu := sync.Mutex{}

	subscribers := map[string]chan mitm.Shot{}
	subscribe := func() (<-chan mitm.Shot, func()) {
		mu.Lock()
		defer mu.Unlock()

		subid := uuid.NewString()
		ch := make(chan mitm.Shot, 100)
		subscribers[subid] = ch

		return ch, func() {
			mu.Lock()
			defer mu.Unlock()

			delete(subscribers, subid)
		}
	}

	/*go func() {
		var shotsToAdd []mitm.Shot
		json.Unmarshal([]byte(shotsJSON), &shotsToAdd)

		for len(shotsToAdd) > 0 {
			time.Sleep(time.Duration(500) * time.Millisecond)
			s := shotsToAdd[0]
			shotsToAdd = shotsToAdd[1:]

			func() {
				mu.Lock()
				defer mu.Unlock()
				shots = append(shots, s)
				for _, ch := range subscribers {
					ch <- s
				}
			}()
		}
	}()*/

	r.Get("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		mu.Lock()
		defer mu.Unlock()

		view := "score"

		if r.URL.Query().Get("view") == "group" {
			view = "group"
		}

		home := components.StartPage(version, assetLoader.Load, shots, view)
		home.Render(r.Context(), w)
	}))

	r.Get("/shootinglog", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		mu.Lock()
		defer mu.Unlock()

		shootingLog := components.ShootingLog(version, assetLoader.Load, shots, true)
		shootingLog.Render(r.Context(), w)
	}))

	r.Route("/components", func(r router.ServeMux) {
		r.Get("/scorecard", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusOK)

			mu.Lock()
			defer mu.Unlock()

			card := components.ScoreCard(shots)
			card.Render(r.Context(), w)
		}))

		r.Get("/shootlog", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusOK)

			mu.Lock()
			defer mu.Unlock()

			card := components.ShootingLogComponent(shots, true)
			card.Render(r.Context(), w)
		}))

		r.Get("/shotgroup", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusOK)

			mu.Lock()
			defer mu.Unlock()

			card := components.ShotGroup(shots)
			card.Render(r.Context(), w)
		}))
	})

	r.Route("/api", func(r router.ServeMux) {
		r.Get("/shots", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body := func() []byte {
				mu.Lock()
				defer mu.Unlock()

				b, _ := json.Marshal(shots)
				return b
			}()

			w.Header().Add("Content-Type", "application/json")
			w.Header().Add(
				"Content-Disposition",
				fmt.Sprintf("attachment; filename=\"siusdata-%d-shots-%s.json\"",
					len(shots),
					time.Now().Format("060102")),
			)
			w.Header().Add("Content-Length", fmt.Sprintf("%d", len(body)))

			w.WriteHeader(http.StatusOK)
			w.Write(body)
		}))

		r.Post("/shots", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			b, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			var shot mitm.Shot
			if err = json.Unmarshal(b, &shot); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			mu.Lock()
			defer mu.Unlock()

			if len(shots) > 0 && shots[len(shots)-1].Idx >= shot.Idx {
				shots = make([]mitm.Shot, 0, 100)
			}

			shots = append(shots, shot)
			for _, ch := range subscribers {
				ch <- shot
			}

			w.WriteHeader(http.StatusOK)
		}))

		r.Get("/sse/{version}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger := logging.GetFromContext(r.Context())

			flusher, ok := w.(http.Flusher)
			if !ok {
				logger.Warn("streaming not supported for this response writer")
				http.Error(w, "unable to start event stream", http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("Connection", "keep-alive")

			const eventFmt string = "event: %s\ndata: %s\n\n"

			logger.Info("comparing versions", "client", r.PathValue("version"), "mine", version)
			waitingForUpgrade := false

			if strings.Compare(r.PathValue("version"), version) != 0 {
				waitingForUpgrade = true

				logger.Info("sending upgrade to client")
				fmt.Fprintf(w, eventFmt, "upgrade2", version)
				flusher.Flush()

				time.Sleep(10 * time.Second)

				logger.Info("sending goodbye to client")
				fmt.Fprintf(w, eventFmt, "goodbye", version)
				flusher.Flush()
			} else {
				logger.Info("sse client successfully connected")
			}

			defer func() { logger.Info("exiting sse handler") }()

			tmr := time.NewTicker(time.Second)
			shotCh, unsubscribe := subscribe()
			defer unsubscribe()

			for {
				select {
				case s := <-shotCh:
					if !waitingForUpgrade {
						fmt.Fprintf(w, eventFmt, "shot", s.String())
					}
				case t := <-tmr.C:
					if !waitingForUpgrade {
						fmt.Fprintf(w, eventFmt, "tick", t.Format(time.RFC3339Nano))
					}
				case <-r.Context().Done():
					logger.Info("sse client connection closed")
					return
				case <-ctx.Done():
					logger.Info("we are shutting down")
					fmt.Fprintf(w, eventFmt, "goodbye", version)
					flusher.Flush()
					return
				}

				flusher.Flush()
			}

		}))
	})

	return nil
}

//const shotsJSON string = `[{"id":"shot-1-18143132-1164327132","idx":1,"score":97,"poi":{"x":-0.0004,"y":-0.003}},{"id":"shot-2-18152569-1164332569","idx":2,"score":82,"poi":{"x":-0.0052,"y":0.0044}},{"id":"shot-3-18161543-1164337543","idx":3,"score":99,"poi":{"x":-0.0004,"y":0.0026}},{"id":"shot-4-18170114-1164342114","idx":4,"score":101,"poi":{"x":0.0016,"y":0.0014}},{"id":"shot-5-18174561-1164346561","idx":5,"score":93,"poi":{"x":0.0019,"y":0.0037}},{"id":"shot-6-18183345-1164351345","idx":6,"score":95,"poi":{"x":-0.0032,"y":0.0018}},{"id":"shot-7-18190781-1164354781","idx":7,"score":88,"poi":{"x":-0.0053,"y":-0.0006}},{"id":"shot-8-18195632-1164359632","idx":8,"score":104,"poi":{"x":0.0004,"y":-0.0014}},{"id":"shot-9-18255024-1164395024","idx":9,"score":102,"poi":{"x":-0.0007,"y":-0.0018}},{"id":"shot-10-18264187-1164400187","idx":10,"score":75,"poi":{"x":0.008,"y":-0.003}},{"id":"shot-11-18273356-1164405356","idx":11,"score":82,"poi":{"x":-0.0068,"y":-0.0002}},{"id":"shot-12-18281885-1164409885","idx":12,"score":73,"poi":{"x":0.0086,"y":0.0033}},{"id":"shot-13-18291216-1164415216","idx":13,"score":99,"poi":{"x":0.0006,"y":-0.0025}},{"id":"shot-14-18300143-1164420143","idx":14,"score":107,"poi":{"x":0.0005,"y":-0.0004}},{"id":"shot-15-18305635-1164425635","idx":15,"score":77,"poi":{"x":-0.0063,"y":-0.005}},{"id":"shot-16-18315874-1164431874","idx":16,"score":83,"poi":{"x":0.006,"y":-0.003}},{"id":"shot-17-18324507-1164436507","idx":17,"score":107,"poi":{"x":-0.0006,"y":0.0002}},{"id":"shot-18-18332987-1164440987","idx":18,"score":88,"poi":{"x":0.0044,"y":-0.0032}},{"id":"shot-19-18342061-1164446061","idx":19,"score":103,"poi":{"x":0.0005,"y":-0.0015}},{"id":"shot-20-18352307-1164452307","idx":20,"score":94,"poi":{"x":-0.0025,"y":-0.0028}},{"id":"shot-21-18361006-1164457006","idx":21,"score":95,"poi":{"x":0.0031,"y":0.002}},{"id":"shot-22-18365783-1164461783","idx":22,"score":95,"poi":{"x":0.0037,"y":-0.0001}},{"id":"shot-23-18375993-1164467993","idx":23,"score":97,"poi":{"x":-0.0005,"y":-0.003}},{"id":"shot-24-18385649-1164473649","idx":24,"score":86,"poi":{"x":-0.0009,"y":0.0059}},{"id":"shot-25-18395287-1164479287","idx":25,"score":89,"poi":{"x":0.0023,"y":-0.0045}},{"id":"shot-26-18404110-1164484110","idx":26,"score":93,"poi":{"x":0.002,"y":-0.0035}},{"id":"shot-27-18412443-1164488443","idx":27,"score":82,"poi":{"x":0.0036,"y":-0.006}},{"id":"shot-28-18420765-1164492765","idx":28,"score":91,"poi":{"x":0.0018,"y":0.0042}},{"id":"shot-29-18425056-1164497056","idx":29,"score":103,"poi":{"x":0.0004,"y":0.0015}},{"id":"shot-30-18433225-1164501225","idx":30,"score":104,"poi":{"x":0.0012,"y":-0.0007}},{"id":"shot-31-18470574-1164522574","idx":31,"score":92,"poi":{"x":-0.0023,"y":-0.0037}},{"id":"shot-32-18480234-1164528234","idx":32,"score":101,"poi":{"x":-0.0003,"y":-0.0021}},{"id":"shot-33-18484603-1164532603","idx":33,"score":80,"poi":{"x":-0.0073,"y":-0.0006}},{"id":"shot-34-18504122-1164544122","idx":34,"score":90,"poi":{"x":0.0048,"y":0}},{"id":"shot-35-18512129-1164548129","idx":35,"score":72,"poi":{"x":-0.0091,"y":-0.0017}},{"id":"shot-36-18520265-1164552265","idx":36,"score":80,"poi":{"x":-0.0073,"y":0.0002}},{"id":"shot-37-18524981-1164556981","idx":37,"score":91,"poi":{"x":-0.0046,"y":0.0006}},{"id":"shot-38-18534065-1164562065","idx":38,"score":91,"poi":{"x":-0.0037,"y":-0.0026}},{"id":"shot-39-18543226-1164567226","idx":39,"score":83,"poi":{"x":-0.0055,"y":-0.0036}},{"id":"shot-40-18551369-1164571369","idx":40,"score":90,"poi":{"x":0.0007,"y":-0.0048}},{"id":"shot-41-18560516-1164576516","idx":41,"score":86,"poi":{"x":0.0008,"y":-0.0059}},{"id":"shot-42-18565275-1164581275","idx":42,"score":101,"poi":{"x":-0.0021,"y":-0.0007}},{"id":"shot-43-18573617-1164585617","idx":43,"score":85,"poi":{"x":0.0057,"y":-0.0023}},{"id":"shot-44-18582918-1164590918","idx":44,"score":83,"poi":{"x":-0.0067,"y":0}},{"id":"shot-45-18591196-1164595196","idx":45,"score":91,"poi":{"x":0.0039,"y":0.0025}},{"id":"shot-46-18594664-1164598664","idx":46,"score":87,"poi":{"x":-0.0054,"y":-0.0017}},{"id":"shot-47-19002454-1164602454","idx":47,"score":96,"poi":{"x":0.0031,"y":-0.001}},{"id":"shot-48-19005440-1164605440","idx":48,"score":81,"poi":{"x":0.006,"y":0.004}},{"id":"shot-49-19013845-1164609845","idx":49,"score":103,"poi":{"x":0.0015,"y":0.0005}},{"id":"shot-50-19021613-1164613613","idx":50,"score":104,"poi":{"x":-0.0009,"y":-0.001}},{"id":"shot-51-19031361-1164619361","idx":51,"score":60,"poi":{"x":0.012,"y":-0.0027}},{"id":"shot-52-19040190-1164624190","idx":52,"score":95,"poi":{"x":-0.0036,"y":0.0003}},{"id":"shot-53-19044649-1164628649","idx":53,"score":79,"poi":{"x":0.007,"y":-0.0029}},{"id":"shot-54-19053182-1164633182","idx":54,"score":99,"poi":{"x":0.0024,"y":0.0008}},{"id":"shot-55-19061930-1164637930","idx":55,"score":88,"poi":{"x":-0.0043,"y":-0.003}},{"id":"shot-56-19065610-1164641610","idx":56,"score":95,"poi":{"x":-0.0004,"y":-0.0035}},{"id":"shot-57-19074107-1164646107","idx":57,"score":80,"poi":{"x":0.007,"y":-0.0023}},{"id":"shot-58-19081899-1164649899","idx":58,"score":99,"poi":{"x":0.0001,"y":0.0027}},{"id":"shot-59-19085393-1164653393","idx":59,"score":57,"poi":{"x":0.0124,"y":-0.0038}},{"id":"shot-60-19094738-1164658738","idx":60,"score":94,"poi":{"x":-0.0039,"y":-0.0005}}]`

//const shotsJSON string = `[{"id":"shot-1-18083336-1207491336","idx":1,"score":97,"poi":{"x":-0.0008,"y":0.003}},{"id":"shot-2-18092314-1207496314","idx":2,"score":99,"poi":{"x":0.0025,"y":-0.0007}},{"id":"shot-3-18101605-1207501605","idx":3,"score":85,"poi":{"x":0.0037,"y":0.0049}},{"id":"shot-4-18110546-1207506546","idx":4,"score":98,"poi":{"x":-0.001,"y":0.0027}},{"id":"shot-5-18115589-1207511589","idx":5,"score":97,"poi":{"x":0.003,"y":-0.0007}},{"id":"shot-6-18131047-1207519047","idx":6,"score":92,"poi":{"x":-0.0032,"y":0.003}},{"id":"shot-7-18141721-1207525721","idx":7,"score":95,"poi":{"x":0,"y":0.0037}},{"id":"shot-8-18151002-1207531002","idx":8,"score":107,"poi":{"x":-0.0007,"y":0}},{"id":"shot-9-18160632-1207536632","idx":9,"score":89,"poi":{"x":-0.0004,"y":0.005}},{"id":"shot-10-18165173-1207541173","idx":10,"score":94,"poi":{"x":-0.0034,"y":0.0018}},{"id":"shot-11-18175451-1207547451","idx":11,"score":86,"poi":{"x":0.0053,"y":0.0023}},{"id":"shot-12-18184453-1207552453","idx":12,"score":106,"poi":{"x":0.0005,"y":-0.0007}},{"id":"shot-13-18194312-1207558312","idx":13,"score":106,"poi":{"x":-0.0005,"y":0.0008}},{"id":"shot-14-18203261-1207563261","idx":14,"score":97,"poi":{"x":0.0013,"y":0.0029}},{"id":"shot-15-18211854-1207567854","idx":15,"score":100,"poi":{"x":-0.0023,"y":-0.0002}},{"id":"shot-16-18220810-1207572810","idx":16,"score":101,"poi":{"x":-0.0016,"y":0.0016}},{"id":"shot-17-18225858-1207577858","idx":17,"score":96,"poi":{"x":-0.0017,"y":0.003}},{"id":"shot-18-18234527-1207582527","idx":18,"score":96,"poi":{"x":0.0032,"y":0.001}},{"id":"shot-19-18243306-1207587306","idx":19,"score":46,"poi":{"x":0.0153,"y":-0.0041}},{"id":"shot-20-18253013-1207593013","idx":20,"score":96,"poi":{"x":0.0026,"y":-0.0019}},{"id":"shot-21-18293834-1207617834","idx":21,"score":87,"poi":{"x":-0.0056,"y":-0.0008}},{"id":"shot-22-18302569-1207622569","idx":22,"score":101,"poi":{"x":0.0004,"y":0.0022}},{"id":"shot-23-18311418-1207627418","idx":23,"score":96,"poi":{"x":-0.0006,"y":-0.0034}},{"id":"shot-24-18315506-1207631506","idx":24,"score":102,"poi":{"x":-0.0013,"y":0.0013}},{"id":"shot-25-18323322-1207635322","idx":25,"score":99,"poi":{"x":0.0007,"y":-0.0026}},{"id":"shot-26-18331735-1207639735","idx":26,"score":107,"poi":{"x":0.0007,"y":-0.0001}},{"id":"shot-27-18340211-1207644211","idx":27,"score":78,"poi":{"x":-0.0074,"y":0.0029}},{"id":"shot-28-18344875-1207648875","idx":28,"score":98,"poi":{"x":-0.0029,"y":-0.0003}},{"id":"shot-29-18353003-1207653003","idx":29,"score":101,"poi":{"x":-0.0012,"y":0.0017}},{"id":"shot-30-18361803-1207657803","idx":30,"score":106,"poi":{"x":-0.0009,"y":-0.0005}},{"id":"shot-31-18365596-1207661596","idx":31,"score":87,"poi":{"x":-0.0044,"y":0.0034}},{"id":"shot-32-18374365-1207666365","idx":32,"score":107,"poi":{"x":0.0001,"y":-0.0007}},{"id":"shot-33-18382143-1207670143","idx":33,"score":89,"poi":{"x":0.0051,"y":-0.0006}},{"id":"shot-34-18390114-1207674114","idx":34,"score":105,"poi":{"x":-0.0004,"y":0.0011}},{"id":"shot-35-18393767-1207677767","idx":35,"score":93,"poi":{"x":0.0007,"y":0.0042}},{"id":"shot-36-18401893-1207681893","idx":36,"score":98,"poi":{"x":0.0027,"y":0.0008}},{"id":"shot-37-18410128-1207686128","idx":37,"score":90,"poi":{"x":-0.0049,"y":0.0007}},{"id":"shot-38-18413672-1207689672","idx":38,"score":93,"poi":{"x":0.0042,"y":-0.0003}},{"id":"shot-39-18421990-1207693990","idx":39,"score":93,"poi":{"x":0.0042,"y":0.0006}},{"id":"shot-40-18430636-1207698636","idx":40,"score":102,"poi":{"x":-0.0014,"y":0.0012}}]`

//const shotsJSON string = `[{"id":"shot-1-18242206-1328546206","idx":1,"score":75,"poi":{"x":-0.0056,"y":0.0066}},{"id":"shot-2-18250616-1328550616","idx":2,"score":106,"poi":{"x":-0.0009,"y":-0.0003}},{"id":"shot-3-18261029-1328557029","idx":3,"score":93,"poi":{"x":0.0008,"y":0.0041}},{"id":"shot-4-18265720-1328561720","idx":4,"score":103,"poi":{"x":-0.0012,"y":0.0012}},{"id":"shot-5-18273477-1328565477","idx":5,"score":91,"poi":{"x":0.0042,"y":0.0017}},{"id":"shot-6-18281725-1328569725","idx":6,"score":105,"poi":{"x":-0.0012,"y":-0.0002}},{"id":"shot-7-18290272-1328574272","idx":7,"score":95,"poi":{"x":0.0026,"y":-0.0025}},{"id":"shot-8-18294696-1328578696","idx":8,"score":102,"poi":{"x":0.0019,"y":0.0004}},{"id":"shot-9-18303715-1328583715","idx":9,"score":106,"poi":{"x":-0.0002,"y":0.0009}},{"id":"shot-10-18311877-1328587877","idx":10,"score":99,"poi":{"x":-0.0012,"y":0.0025}},{"id":"shot-11-18323010-1328595010","idx":11,"score":100,"poi":{"x":0.0013,"y":0.0019}},{"id":"shot-12-18331861-1328599861","idx":12,"score":99,"poi":{"x":0.0025,"y":-0.0002}},{"id":"shot-13-18335628-1328603628","idx":13,"score":77,"poi":{"x":0.0072,"y":-0.0038}},{"id":"shot-14-18344174-1328608174","idx":14,"score":101,"poi":{"x":-0.0018,"y":0.0013}},{"id":"shot-15-18352963-1328612963","idx":15,"score":61,"poi":{"x":0.012,"y":0.0015}},{"id":"shot-16-18361351-1328617351","idx":16,"score":98,"poi":{"x":0.0024,"y":-0.0015}},{"id":"shot-17-18365430-1328621430","idx":17,"score":95,"poi":{"x":0.0034,"y":-0.0013}},{"id":"shot-18-18374250-1328626250","idx":18,"score":84,"poi":{"x":0.0061,"y":0.0016}},{"id":"shot-19-18383144-1328631144","idx":19,"score":80,"poi":{"x":0.0073,"y":0.0018}},{"id":"shot-20-18391263-1328635263","idx":20,"score":92,"poi":{"x":0.0045,"y":0.0001}},{"id":"shot-21-18401484-1328641484","idx":21,"score":85,"poi":{"x":0.0061,"y":-0.0008}},{"id":"shot-22-18405902-1328645902","idx":22,"score":102,"poi":{"x":-0.0019,"y":-0.0007}},{"id":"shot-23-18414114-1328650114","idx":23,"score":87,"poi":{"x":0.0002,"y":0.0057}},{"id":"shot-24-18422041-1328654041","idx":24,"score":99,"poi":{"x":-0.0025,"y":0}},{"id":"shot-25-18430537-1328658537","idx":25,"score":91,"poi":{"x":0.0046,"y":0.0002}},{"id":"shot-26-18434337-1328662337","idx":26,"score":93,"poi":{"x":-0.0018,"y":0.0037}},{"id":"shot-27-18444098-1328668098","idx":27,"score":95,"poi":{"x":-0.0021,"y":0.0028}},{"id":"shot-28-18452955-1328672955","idx":28,"score":102,"poi":{"x":0.002,"y":-0.0001}},{"id":"shot-29-18461005-1328677005","idx":29,"score":101,"poi":{"x":0.002,"y":0.0003}},{"id":"shot-30-18470864-1328682864","idx":30,"score":85,"poi":{"x":-0.0028,"y":0.0053}},{"id":"shot-31-18474967-1328686967","idx":31,"score":88,"poi":{"x":-0.0053,"y":0.0009}},{"id":"shot-32-18483703-1328691703","idx":32,"score":99,"poi":{"x":-0.0024,"y":0.0009}},{"id":"shot-33-18492723-1328696723","idx":33,"score":101,"poi":{"x":-0.0015,"y":-0.0015}},{"id":"shot-34-18501029-1328701029","idx":34,"score":92,"poi":{"x":-0.0006,"y":0.0043}},{"id":"shot-35-18504702-1328704702","idx":35,"score":99,"poi":{"x":-0.0027,"y":0.0004}},{"id":"shot-36-18512171-1328708171","idx":36,"score":105,"poi":{"x":-0.0009,"y":0.0007}},{"id":"shot-37-18520939-1328712939","idx":37,"score":75,"poi":{"x":-0.0081,"y":-0.003}},{"id":"shot-38-18525103-1328717103","idx":38,"score":67,"poi":{"x":0.0103,"y":0.0024}},{"id":"shot-39-18533014-1328721014","idx":39,"score":88,"poi":{"x":0.0054,"y":-0.0008}},{"id":"shot-40-18541494-1328725494","idx":40,"score":103,"poi":{"x":-0.0016,"y":-0.0003}}]`

// 2025-09-11
//const shotsJSON string = `[{"id":"shot-1-17310952-2174946952","idx":1,"score":89,"poi":{"x":-0.0047,"y":0.0021}},{"id":"shot-2-17314624-2174950624","idx":2,"score":67,"poi":{"x":-0.0105,"y":0.0015}},{"id":"shot-3-17323796-2174955796","idx":3,"score":92,"poi":{"x":-0.0016,"y":0.0042}},{"id":"shot-4-17333071-2174961071","idx":4,"score":84,"poi":{"x":-0.0051,"y":0.0037}},{"id":"shot-5-17343827-2174967827","idx":5,"score":81,"poi":{"x":-0.0052,"y":0.0049}},{"id":"shot-6-17353187-2174973187","idx":6,"score":95,"poi":{"x":0.0021,"y":0.003}},{"id":"shot-7-17361937-2174977937","idx":7,"score":89,"poi":{"x":0.0026,"y":0.0043}},{"id":"shot-8-17370331-2174982331","idx":8,"score":80,"poi":{"x":-0.0073,"y":0.0007}},{"id":"shot-9-17374450-2174986450","idx":9,"score":105,"poi":{"x":0.0002,"y":0.0012}},{"id":"shot-10-17383261-2174991261","idx":10,"score":95,"poi":{"x":0.0025,"y":0.0025}},{"id":"shot-11-17392174-2174996174","idx":11,"score":97,"poi":{"x":0.0011,"y":0.0029}},{"id":"shot-12-17395139-2174999139","idx":12,"score":61,"poi":{"x":-0.0107,"y":0.0055}},{"id":"shot-13-17403537-2175003537","idx":13,"score":83,"poi":{"x":-0.0057,"y":0.0033}},{"id":"shot-14-17412334-2175008334","idx":14,"score":78,"poi":{"x":-0.0036,"y":0.007}},{"id":"shot-15-17420072-2175012072","idx":15,"score":89,"poi":{"x":-0.003,"y":0.0043}},{"id":"shot-16-17432821-2175020821","idx":16,"score":49,"poi":{"x":0.0152,"y":-0.0001}},{"id":"shot-17-17441121-2175025121","idx":17,"score":83,"poi":{"x":0.006,"y":0.0029}},{"id":"shot-18-17450634-2175030634","idx":18,"score":87,"poi":{"x":0.005,"y":0.0025}},{"id":"shot-19-17454889-2175034889","idx":19,"score":83,"poi":{"x":-0.0032,"y":0.0057}},{"id":"shot-20-17465364-2175041364","idx":20,"score":78,"poi":{"x":-0.0068,"y":0.0041}},{"id":"shot-21-17473867-2175045867","idx":21,"score":87,"poi":{"x":0.0043,"y":0.0037}},{"id":"shot-22-17481650-2175049650","idx":22,"score":80,"poi":{"x":-0.0061,"y":0.0042}},{"id":"shot-23-17485952-2175053952","idx":23,"score":99,"poi":{"x":0.0001,"y":0.0025}},{"id":"shot-24-17494624-2175058624","idx":24,"score":105,"poi":{"x":-0.0001,"y":-0.0011}},{"id":"shot-25-17503234-2175063234","idx":25,"score":104,"poi":{"x":0.0013,"y":-0.0003}},{"id":"shot-26-17512030-2175068030","idx":26,"score":89,"poi":{"x":0.0042,"y":0.0029}},{"id":"shot-27-17521118-2175073118","idx":27,"score":85,"poi":{"x":-0.0054,"y":0.0031}},{"id":"shot-28-17530656-2175078656","idx":28,"score":95,"poi":{"x":-0.0033,"y":0.0015}},{"id":"shot-29-17540713-2175084713","idx":29,"score":94,"poi":{"x":-0.0037,"y":0.0003}},{"id":"shot-30-17545993-2175089993","idx":30,"score":93,"poi":{"x":-0.0041,"y":-0.0005}},{"id":"shot-31-17553830-2175093830","idx":31,"score":96,"poi":{"x":-0.0031,"y":-0.0012}},{"id":"shot-32-17562422-2175098422","idx":32,"score":105,"poi":{"x":-0.0011,"y":0.0002}},{"id":"shot-33-17570797-2175102797","idx":33,"score":73,"poi":{"x":-0.009,"y":-0.002}},{"id":"shot-34-17575479-2175107479","idx":34,"score":75,"poi":{"x":-0.0085,"y":0.002}},{"id":"shot-35-17583665-2175111665","idx":35,"score":90,"poi":{"x":-0.0046,"y":0.0014}},{"id":"shot-36-17592424-2175116424","idx":36,"score":86,"poi":{"x":-0.0057,"y":0.0018}},{"id":"shot-37-18000654-2175120654","idx":37,"score":97,"poi":{"x":0.0002,"y":-0.0032}},{"id":"shot-38-18005004-2175125004","idx":38,"score":72,"poi":{"x":0.0093,"y":-0.0007}},{"id":"shot-39-18013248-2175129248","idx":39,"score":100,"poi":{"x":-0.0011,"y":0.0021}},{"id":"shot-40-18021554-2175133554","idx":40,"score":102,"poi":{"x":-0.0015,"y":-0.0013}}]`

// 2025-09-16
//const shotsJSON string = `[{"id":"shot-1-17450399-2235510399","idx":1,"score":95,"poi":{"x":0.0035,"y":0.0003}},{"id":"shot-2-17454988-2235514988","idx":2,"score":85,"poi":{"x":0.0059,"y":0.0012}},{"id":"shot-3-17463542-2235519542","idx":3,"score":89,"poi":{"x":-0.0038,"y":0.0034}},{"id":"shot-4-17472183-2235524183","idx":4,"score":97,"poi":{"x":-0.0026,"y":0.0019}},{"id":"shot-5-17481186-2235529186","idx":5,"score":75,"poi":{"x":-0.0075,"y":0.0044}},{"id":"shot-6-17490823-2235534823","idx":6,"score":101,"poi":{"x":-0.0002,"y":-0.002}},{"id":"shot-7-17500072-2235540072","idx":7,"score":74,"poi":{"x":0.0083,"y":0.0027}},{"id":"shot-8-17505402-2235545402","idx":8,"score":47,"poi":{"x":-0.0156,"y":-0.0005}},{"id":"shot-9-17513854-2235549854","idx":9,"score":77,"poi":{"x":0.0045,"y":0.0069}},{"id":"shot-10-17530471-2235558471","idx":10,"score":81,"poi":{"x":-0.007,"y":0.0014}},{"id":"shot-11-17534766-2235562766","idx":11,"score":93,"poi":{"x":-0.004,"y":0.0012}},{"id":"shot-12-17545190-2235569190","idx":12,"score":101,"poi":{"x":0.0003,"y":0.002}},{"id":"shot-13-17555851-2235575851","idx":13,"score":102,"poi":{"x":-0.0019,"y":-0.0005}},{"id":"shot-14-17564861-2235580861","idx":14,"score":86,"poi":{"x":-0.0058,"y":-0.0011}},{"id":"shot-15-17574180-2235586180","idx":15,"score":98,"poi":{"x":-0.0029,"y":0.0002}},{"id":"shot-16-17584158-2235592158","idx":16,"score":98,"poi":{"x":-0.0024,"y":-0.0014}},{"id":"shot-17-17594431-2235598431","idx":17,"score":99,"poi":{"x":-0.002,"y":0.0018}},{"id":"shot-18-18003680-2235603680","idx":18,"score":104,"poi":{"x":0.0009,"y":0.0009}},{"id":"shot-19-18013303-2235609303","idx":19,"score":108,"poi":{"x":0.0004,"y":-0.0003}},{"id":"shot-20-18022968-2235614968","idx":20,"score":86,"poi":{"x":-0.0053,"y":-0.0026}},{"id":"shot-21-18032778-2235620778","idx":21,"score":103,"poi":{"x":0.0001,"y":0.0016}},{"id":"shot-22-18042566-2235626566","idx":22,"score":86,"poi":{"x":0.0047,"y":0.0033}},{"id":"shot-23-18051654-2235631654","idx":23,"score":94,"poi":{"x":-0.0021,"y":-0.0033}},{"id":"shot-24-18060761-2235636761","idx":24,"score":106,"poi":{"x":-0.0004,"y":-0.0008}},{"id":"shot-25-18065128-2235641128","idx":25,"score":97,"poi":{"x":0.0015,"y":0.0026}},{"id":"shot-26-18073984-2235645984","idx":26,"score":74,"poi":{"x":0.009,"y":0.0001}},{"id":"shot-27-18082911-2235650911","idx":27,"score":83,"poi":{"x":-0.0054,"y":0.004}},{"id":"shot-28-18091959-2235655959","idx":28,"score":101,"poi":{"x":0.0009,"y":0.002}},{"id":"shot-29-18100190-2235660190","idx":29,"score":83,"poi":{"x":-0.0023,"y":0.0061}},{"id":"shot-30-18105070-2235665070","idx":30,"score":87,"poi":{"x":0.0001,"y":0.0056}},{"id":"shot-31-18113467-2235669467","idx":31,"score":88,"poi":{"x":0.0014,"y":0.0052}},{"id":"shot-32-18121907-2235673907","idx":32,"score":73,"poi":{"x":-0.0092,"y":0.0002}},{"id":"shot-33-18130064-2235678064","idx":33,"score":86,"poi":{"x":0.0057,"y":-0.0005}},{"id":"shot-34-18134611-2235682611","idx":34,"score":99,"poi":{"x":0.0002,"y":0.0027}},{"id":"shot-35-18142901-2235686901","idx":35,"score":73,"poi":{"x":-0.0081,"y":0.0044}},{"id":"shot-36-18150856-2235690856","idx":36,"score":78,"poi":{"x":-0.0011,"y":0.0077}},{"id":"shot-37-18155356-2235695356","idx":37,"score":83,"poi":{"x":0.0064,"y":0.0015}},{"id":"shot-38-18163527-2235699527","idx":38,"score":104,"poi":{"x":0.0009,"y":-0.001}},{"id":"shot-39-18171102-2235703102","idx":39,"score":107,"poi":{"x":0,"y":-0.0005}},{"id":"shot-40-18175398-2235707398","idx":40,"score":81,"poi":{"x":0.006,"y":0.0041}}]`

// 2025-09-18
//const shotsJSON string = `[{"id":"shot-1-18050568-2252910568","idx":1,"score":79,"poi":{"x":-0.0076,"y":-0.0014}},{"id":"shot-2-18060453-2252916453","idx":2,"score":91,"poi":{"x":0.0042,"y":-0.0021}},{"id":"shot-3-18065341-2252921341","idx":3,"score":95,"poi":{"x":-0.0037,"y":0}},{"id":"shot-4-18074179-2252926179","idx":4,"score":103,"poi":{"x":-0.0014,"y":0.0007}},{"id":"shot-5-18083955-2252931955","idx":5,"score":106,"poi":{"x":0.0002,"y":-0.0007}},{"id":"shot-6-18092913-2252936913","idx":6,"score":101,"poi":{"x":0.002,"y":-0.0005}},{"id":"shot-7-18101819-2252941819","idx":7,"score":97,"poi":{"x":-0.0026,"y":-0.0019}},{"id":"shot-8-18110026-2252946026","idx":8,"score":88,"poi":{"x":0.0004,"y":0.0053}},{"id":"shot-9-18115159-2252951159","idx":9,"score":103,"poi":{"x":-0.0016,"y":-0.0007}},{"id":"shot-10-18123869-2252955869","idx":10,"score":104,"poi":{"x":-0.0005,"y":-0.0013}},{"id":"shot-11-18295253-2253059253","idx":11,"score":78,"poi":{"x":0.0076,"y":0.0016}},{"id":"shot-12-18305932-2253065932","idx":12,"score":101,"poi":{"x":0.0021,"y":0.0001}},{"id":"shot-13-18320338-2253072338","idx":13,"score":95,"poi":{"x":0.0024,"y":-0.0029}},{"id":"shot-14-18325865-2253077865","idx":14,"score":95,"poi":{"x":0.0033,"y":0.0014}},{"id":"shot-15-18335532-2253083532","idx":15,"score":86,"poi":{"x":-0.0057,"y":-0.001}},{"id":"shot-16-18343987-2253087987","idx":16,"score":98,"poi":{"x":0.0029,"y":0.0007}},{"id":"shot-17-18353028-2253093028","idx":17,"score":107,"poi":{"x":0.0006,"y":0.0004}},{"id":"shot-18-18364123-2253100123","idx":18,"score":93,"poi":{"x":0.0026,"y":-0.0033}},{"id":"shot-19-18372852-2253104852","idx":19,"score":107,"poi":{"x":-0.0003,"y":-0.0005}},{"id":"shot-20-18383282-2253111282","idx":20,"score":87,"poi":{"x":0.0036,"y":-0.0043}},{"id":"shot-21-18393762-2253117762","idx":21,"score":107,"poi":{"x":-0.0001,"y":-0.0007}},{"id":"shot-22-18403993-2253123993","idx":22,"score":101,"poi":{"x":0.0021,"y":0.0003}},{"id":"shot-23-18413652-2253129652","idx":23,"score":108,"poi":{"x":-0.0002,"y":-0.0003}},{"id":"shot-24-18423001-2253135001","idx":24,"score":93,"poi":{"x":0.0032,"y":0.0027}},{"id":"shot-25-18432880-2253140880","idx":25,"score":93,"poi":{"x":-0.0018,"y":0.0038}},{"id":"shot-26-18441838-2253145838","idx":26,"score":95,"poi":{"x":0.0031,"y":-0.0018}},{"id":"shot-27-18451673-2253151673","idx":27,"score":73,"poi":{"x":0.0071,"y":0.0056}},{"id":"shot-28-18461950-2253157950","idx":28,"score":92,"poi":{"x":-0.0043,"y":-0.0004}},{"id":"shot-29-18470862-2253162862","idx":29,"score":87,"poi":{"x":-0.0057,"y":0.0005}},{"id":"shot-30-18480749-2253168749","idx":30,"score":90,"poi":{"x":-0.0048,"y":-0.0012}},{"id":"shot-31-18490163-2253174163","idx":31,"score":95,"poi":{"x":-0.0021,"y":-0.003}},{"id":"shot-32-18500046-2253180046","idx":32,"score":88,"poi":{"x":-0.0055,"y":0.0005}},{"id":"shot-33-18505207-2253185207","idx":33,"score":107,"poi":{"x":-0.0002,"y":0.0007}},{"id":"shot-34-18514817-2253190817","idx":34,"score":81,"poi":{"x":-0.0071,"y":0}},{"id":"shot-35-18524109-2253196109","idx":35,"score":88,"poi":{"x":-0.0031,"y":0.0043}},{"id":"shot-36-18532414-2253200414","idx":36,"score":98,"poi":{"x":0.0029,"y":-0.0002}},{"id":"shot-37-18541131-2253205131","idx":37,"score":91,"poi":{"x":-0.0044,"y":0.0018}},{"id":"shot-38-18545837-2253209837","idx":38,"score":86,"poi":{"x":-0.0058,"y":0.0005}},{"id":"shot-39-18555961-2253215961","idx":39,"score":85,"poi":{"x":-0.0058,"y":0.0023}},{"id":"shot-40-18570752-2253222752","idx":40,"score":100,"poi":{"x":-0.0022,"y":0.0009}}]`

// 2025-09-23
//const shotsJSON string = `[{"id":"shot-1-17342286-2295926286","idx":1,"score":101,"poi":{"x":-0.001,"y":0.0019}},{"id":"shot-2-17351014-2295931014","idx":2,"score":79,"poi":{"x":0.0026,"y":0.0071}},{"id":"shot-3-17355388-2295935388","idx":3,"score":78,"poi":{"x":-0.0075,"y":0.0023}},{"id":"shot-4-17363960-2295939960","idx":4,"score":103,"poi":{"x":-0.0003,"y":-0.0015}},{"id":"shot-5-17373071-2295945071","idx":5,"score":108,"poi":{"x":0.0003,"y":0.0002}},{"id":"shot-6-17381972-2295949972","idx":6,"score":107,"poi":{"x":0.0003,"y":0.0005}},{"id":"shot-7-17392134-2295956134","idx":7,"score":96,"poi":{"x":0.0034,"y":-0.0004}},{"id":"shot-8-17400424-2295960424","idx":8,"score":81,"poi":{"x":-0.0066,"y":0.0024}},{"id":"shot-9-17404789-2295964789","idx":9,"score":94,"poi":{"x":-0.0032,"y":0.0024}},{"id":"shot-10-17414244-2295970244","idx":10,"score":98,"poi":{"x":0.0026,"y":0.0014}},{"id":"shot-11-17424725-2295976725","idx":11,"score":91,"poi":{"x":-0.0024,"y":0.0039}},{"id":"shot-12-17433102-2295981102","idx":12,"score":89,"poi":{"x":-0.0034,"y":0.0037}},{"id":"shot-13-17444288-2295988288","idx":13,"score":93,"poi":{"x":-0.002,"y":0.0036}},{"id":"shot-14-17453787-2295993787","idx":14,"score":64,"poi":{"x":0.0114,"y":-0.0007}},{"id":"shot-15-17462206-2295998206","idx":15,"score":97,"poi":{"x":0.0016,"y":-0.0027}},{"id":"shot-16-17470188-2296002188","idx":16,"score":96,"poi":{"x":-0.0003,"y":-0.0033}},{"id":"shot-17-17474743-2296006743","idx":17,"score":101,"poi":{"x":-0.002,"y":0.0009}},{"id":"shot-18-17483187-2296011187","idx":18,"score":101,"poi":{"x":-0.0018,"y":0.0012}},{"id":"shot-19-17491913-2296015913","idx":19,"score":97,"poi":{"x":0.0027,"y":0.0018}},{"id":"shot-20-17495949-2296019949","idx":20,"score":103,"poi":{"x":-0.0006,"y":-0.0016}},{"id":"shot-21-17504457-2296024457","idx":21,"score":96,"poi":{"x":-0.0032,"y":0.0007}},{"id":"shot-22-17513407-2296029407","idx":22,"score":92,"poi":{"x":0.0039,"y":-0.0019}},{"id":"shot-23-17522712-2296034712","idx":23,"score":86,"poi":{"x":-0.0006,"y":-0.0059}},{"id":"shot-24-17532012-2296040012","idx":24,"score":86,"poi":{"x":0.0055,"y":0.0022}},{"id":"shot-25-17540381-2296044381","idx":25,"score":87,"poi":{"x":-0.0057,"y":-0.0004}},{"id":"shot-26-17551527-2296051527","idx":26,"score":75,"poi":{"x":-0.006,"y":-0.0064}},{"id":"shot-27-17562157-2296058157","idx":27,"score":99,"poi":{"x":-0.0004,"y":-0.0026}},{"id":"shot-28-17571531-2296063531","idx":28,"score":77,"poi":{"x":-0.0071,"y":-0.004}},{"id":"shot-29-17581945-2296069945","idx":29,"score":86,"poi":{"x":0.0033,"y":-0.005}},{"id":"shot-30-17591431-2296075431","idx":30,"score":82,"poi":{"x":-0.0037,"y":-0.0059}},{"id":"shot-31-18002190-2296082190","idx":31,"score":103,"poi":{"x":0.0012,"y":0.001}},{"id":"shot-32-18011106-2296087106","idx":32,"score":78,"poi":{"x":-0.0067,"y":-0.0041}},{"id":"shot-33-18015650-2296091650","idx":33,"score":92,"poi":{"x":-0.0043,"y":-0.0011}},{"id":"shot-34-18023259-2296095259","idx":34,"score":93,"poi":{"x":-0.0035,"y":-0.0025}},{"id":"shot-35-18031058-2296099058","idx":35,"score":104,"poi":{"x":0.0013,"y":-0.0001}},{"id":"shot-36-18035885-2296103885","idx":36,"score":99,"poi":{"x":0.0025,"y":-0.0006}},{"id":"shot-37-18045226-2296109226","idx":37,"score":86,"poi":{"x":0.0013,"y":-0.0057}},{"id":"shot-38-18053568-2296113568","idx":38,"score":81,"poi":{"x":-0.0072,"y":0.0005}},{"id":"shot-39-18063172-2296119172","idx":39,"score":92,"poi":{"x":0.0015,"y":-0.0042}},{"id":"shot-40-18071999-2296123999","idx":40,"score":92,"poi":{"x":0.0017,"y":-0.0039}}]`

// 2025-09-25
//const shotsJSON string = `[{"id":"shot-1-17544238-2313328238","idx":1,"score":92,"poi":{"x":0.0032,"y":-0.0029}},{"id":"shot-2-17554026-2313334026","idx":2,"score":84,"poi":{"x":0.0063,"y":-0.0008}},{"id":"shot-3-17563357-2313339357","idx":3,"score":71,"poi":{"x":0.0091,"y":-0.003}},{"id":"shot-4-17571979-2313343979","idx":4,"score":86,"poi":{"x":-0.0053,"y":0.0022}},{"id":"shot-5-17581242-2313349242","idx":5,"score":103,"poi":{"x":0.0014,"y":0.0009}},{"id":"shot-6-17590817-2313354817","idx":6,"score":99,"poi":{"x":0.0022,"y":-0.0012}},{"id":"shot-7-18001065-2313361065","idx":7,"score":97,"poi":{"x":-0.0032,"y":0.0003}},{"id":"shot-8-18011083-2313367083","idx":8,"score":67,"poi":{"x":0.0103,"y":-0.0026}},{"id":"shot-9-18021383-2313373383","idx":9,"score":99,"poi":{"x":0.0002,"y":0.0025}},{"id":"shot-10-18030891-2313378891","idx":10,"score":99,"poi":{"x":0.0026,"y":0.0001}},{"id":"shot-11-18040288-2313384288","idx":11,"score":96,"poi":{"x":-0.0031,"y":0.0015}},{"id":"shot-12-18050199-2313390199","idx":12,"score":103,"poi":{"x":-0.0014,"y":-0.0009}},{"id":"shot-13-18060308-2313396308","idx":13,"score":79,"poi":{"x":0.0036,"y":0.0067}},{"id":"shot-14-18065402-2313401402","idx":14,"score":94,"poi":{"x":-0.0036,"y":-0.0013}},{"id":"shot-15-18074065-2313406065","idx":15,"score":104,"poi":{"x":0.0009,"y":-0.001}},{"id":"shot-16-18083066-2313411066","idx":16,"score":99,"poi":{"x":-0.001,"y":-0.0025}},{"id":"shot-17-18092126-2313416126","idx":17,"score":99,"poi":{"x":0.0026,"y":0.0003}},{"id":"shot-18-18101474-2313421474","idx":18,"score":75,"poi":{"x":0.0076,"y":-0.004}},{"id":"shot-19-18111290-2313427290","idx":19,"score":107,"poi":{"x":0.0005,"y":-0.0002}},{"id":"shot-20-18120380-2313432380","idx":20,"score":104,"poi":{"x":0.0008,"y":0.0011}},{"id":"shot-21-18130231-2313438231","idx":21,"score":100,"poi":{"x":0,"y":0.0023}},{"id":"shot-22-18135014-2313443014","idx":22,"score":100,"poi":{"x":0.0022,"y":0.0008}},{"id":"shot-23-18144395-2313448395","idx":23,"score":107,"poi":{"x":0.0005,"y":-0.0004}},{"id":"shot-24-18152801-2313452801","idx":24,"score":95,"poi":{"x":0.003,"y":-0.002}},{"id":"shot-25-18161392-2313457392","idx":25,"score":98,"poi":{"x":0.0014,"y":0.0025}},{"id":"shot-26-18170997-2313462997","idx":26,"score":103,"poi":{"x":0.0015,"y":-0.0003}},{"id":"shot-27-18180588-2313468588","idx":27,"score":97,"poi":{"x":0.0024,"y":0.0021}},{"id":"shot-28-18185283-2313473283","idx":28,"score":100,"poi":{"x":-0.002,"y":0.0011}},{"id":"shot-29-18193235-2313477235","idx":29,"score":83,"poi":{"x":-0.0065,"y":0.0011}},{"id":"shot-30-18202059-2313482059","idx":30,"score":95,"poi":{"x":0.0014,"y":0.0034}},{"id":"shot-31-18210693-2313486693","idx":31,"score":100,"poi":{"x":0.0022,"y":0.0003}},{"id":"shot-32-18215412-2313491412","idx":32,"score":99,"poi":{"x":-0.002,"y":-0.0018}},{"id":"shot-33-18224187-2313496187","idx":33,"score":86,"poi":{"x":0.0033,"y":-0.0048}},{"id":"shot-34-18232746-2313500746","idx":34,"score":83,"poi":{"x":0.006,"y":-0.0029}},{"id":"shot-35-18241267-2313505267","idx":35,"score":100,"poi":{"x":-0.0018,"y":-0.0017}},{"id":"shot-36-18250328-2313510328","idx":36,"score":70,"poi":{"x":-0.0087,"y":-0.0047}},{"id":"shot-37-18254582-2313514582","idx":37,"score":98,"poi":{"x":-0.0025,"y":-0.0015}},{"id":"shot-38-18263494-2313519494","idx":38,"score":108,"poi":{"x":0.0004,"y":-0.0003}},{"id":"shot-39-18271950-2313523950","idx":39,"score":79,"poi":{"x":-0.0072,"y":-0.0027}},{"id":"shot-40-18281538-2313529538","idx":40,"score":96,"poi":{"x":-0.0027,"y":-0.0022}}]`

// 2025-10-07
//const shotsJSON string = `[{"id":"shot-1-18093710-2417097710","idx":1,"score":87,"poi":{"x":0.0056,"y":-0.0004}},{"id":"shot-2-18103777-2417103777","idx":2,"score":92,"poi":{"x":-0.0028,"y":0.0034}},{"id":"shot-3-18113295-2417109295","idx":3,"score":99,"poi":{"x":0.0026,"y":0.0005}},{"id":"shot-4-18121688-2417113688","idx":4,"score":85,"poi":{"x":0.0062,"y":-0.0002}},{"id":"shot-5-18130784-2417118784","idx":5,"score":95,"poi":{"x":-0.003,"y":-0.002}},{"id":"shot-6-18135593-2417123593","idx":6,"score":102,"poi":{"x":0.0016,"y":0.0008}},{"id":"shot-7-18144492-2417128492","idx":7,"score":99,"poi":{"x":0.0014,"y":0.0022}},{"id":"shot-8-18152678-2417132678","idx":8,"score":106,"poi":{"x":0,"y":0.0008}},{"id":"shot-9-18194815-2417158815","idx":9,"score":90,"poi":{"x":0.0046,"y":0.0012}},{"id":"shot-10-18205038-2417165038","idx":10,"score":84,"poi":{"x":-0.0062,"y":-0.0007}},{"id":"shot-11-18215561-2417171561","idx":11,"score":96,"poi":{"x":-0.0005,"y":-0.0034}},{"id":"shot-12-18224437-2417176437","idx":12,"score":85,"poi":{"x":0.0061,"y":0.0008}},{"id":"shot-13-18233224-2417181224","idx":13,"score":101,"poi":{"x":0.0021,"y":-0.0006}},{"id":"shot-14-18243674-2417187674","idx":14,"score":87,"poi":{"x":0.0055,"y":0.0014}},{"id":"shot-15-18252995-2417192995","idx":15,"score":91,"poi":{"x":0.0033,"y":-0.0034}},{"id":"shot-16-18261706-2417197706","idx":16,"score":88,"poi":{"x":-0.002,"y":-0.0049}},{"id":"shot-17-18270987-2417202987","idx":17,"score":102,"poi":{"x":0.0018,"y":-0.0002}},{"id":"shot-18-18284432-2417212432","idx":18,"score":95,"poi":{"x":0.0012,"y":-0.0033}},{"id":"shot-19-18292102-2417216102","idx":19,"score":88,"poi":{"x":-0.0055,"y":0.0002}},{"id":"shot-20-18295222-2417219222","idx":20,"score":100,"poi":{"x":0.0016,"y":-0.0017}},{"id":"shot-21-18303119-2417223119","idx":21,"score":89,"poi":{"x":-0.005,"y":-0.0003}},{"id":"shot-22-18311191-2417227191","idx":22,"score":104,"poi":{"x":-0.0002,"y":-0.0014}},{"id":"shot-23-18345135-2417249135","idx":23,"score":84,"poi":{"x":0.0048,"y":-0.0042}},{"id":"shot-24-18354404-2417254404","idx":24,"score":85,"poi":{"x":0.0051,"y":-0.0034}},{"id":"shot-25-18363651-2417259651","idx":25,"score":96,"poi":{"x":-0.0032,"y":0.0007}},{"id":"shot-26-18372095-2417264095","idx":26,"score":82,"poi":{"x":0.0057,"y":-0.0037}},{"id":"shot-27-18381886-2417269886","idx":27,"score":89,"poi":{"x":0.0047,"y":-0.0021}},{"id":"shot-28-18390330-2417274330","idx":28,"score":83,"poi":{"x":0.0029,"y":-0.006}},{"id":"shot-29-18395189-2417279189","idx":29,"score":97,"poi":{"x":-0.003,"y":-0.0011}},{"id":"shot-30-18403334-2417283334","idx":30,"score":101,"poi":{"x":-0.0015,"y":-0.0014}},{"id":"shot-31-18413054-2417289054","idx":31,"score":93,"poi":{"x":-0.0041,"y":0.0003}},{"id":"shot-32-18421050-2417293050","idx":32,"score":77,"poi":{"x":-0.008,"y":0.0018}},{"id":"shot-33-18425384-2417297384","idx":33,"score":107,"poi":{"x":0.0005,"y":0}},{"id":"shot-34-18435889-2417303889","idx":34,"score":71,"poi":{"x":0.009,"y":-0.003}},{"id":"shot-35-18444025-2417308025","idx":35,"score":98,"poi":{"x":-0.0012,"y":-0.0025}},{"id":"shot-36-18452740-2417312740","idx":36,"score":89,"poi":{"x":0.0044,"y":-0.0024}},{"id":"shot-37-18462350-2417318350","idx":37,"score":107,"poi":{"x":0.0002,"y":0.0005}},{"id":"shot-38-18471443-2417323443","idx":38,"score":102,"poi":{"x":0.0017,"y":0.0004}},{"id":"shot-39-18480214-2417328214","idx":39,"score":104,"poi":{"x":-0.0001,"y":0.0013}},{"id":"shot-40-18485412-2417333412","idx":40,"score":97,"poi":{"x":0.003,"y":-0.0003}}]`

// 2025-10-09
//const shotsJSON string = `[{"id":"shot-1-18072231-2434364231","idx":1,"score":104,"poi":{"x":0.0012,"y":-0.0009}},{"id":"shot-2-18082423-2434370423","idx":2,"score":93,"poi":{"x":-0.0011,"y":0.0041}},{"id":"shot-3-18092078-2434376078","idx":3,"score":99,"poi":{"x":0.0018,"y":0.0018}},{"id":"shot-4-18102088-2434382088","idx":4,"score":101,"poi":{"x":0.0022,"y":0.0001}},{"id":"shot-5-18113665-2434389665","idx":5,"score":94,"poi":{"x":-0.003,"y":0.0024}},{"id":"shot-6-18122938-2434394938","idx":6,"score":99,"poi":{"x":0.0022,"y":-0.0016}},{"id":"shot-7-18133639-2434401639","idx":7,"score":101,"poi":{"x":0.0001,"y":0.0022}},{"id":"shot-8-18144509-2434408509","idx":8,"score":102,"poi":{"x":0.0015,"y":-0.0013}},{"id":"shot-9-18154594-2434414594","idx":9,"score":74,"poi":{"x":-0.0086,"y":0.0027}},{"id":"shot-10-18165558-2434421558","idx":10,"score":84,"poi":{"x":0.0057,"y":-0.0027}},{"id":"shot-11-18174097-2434426097","idx":11,"score":90,"poi":{"x":-0.0035,"y":0.0033}},{"id":"shot-12-18182467-2434430467","idx":12,"score":91,"poi":{"x":0.0033,"y":-0.0032}},{"id":"shot-13-18190679-2434434679","idx":13,"score":66,"poi":{"x":-0.0109,"y":-0.0013}},{"id":"shot-14-18195279-2434439279","idx":14,"score":94,"poi":{"x":0.004,"y":0.0001}},{"id":"shot-15-18213041-2434449041","idx":15,"score":89,"poi":{"x":0.0049,"y":-0.0011}},{"id":"shot-16-18221439-2434453439","idx":16,"score":99,"poi":{"x":0.0026,"y":0.0002}},{"id":"shot-17-18230818-2434458818","idx":17,"score":108,"poi":{"x":0,"y":0.0003}},{"id":"shot-18-18235660-2434463660","idx":18,"score":90,"poi":{"x":0.0049,"y":0.0005}},{"id":"shot-19-18245385-2434469385","idx":19,"score":95,"poi":{"x":-0.0023,"y":0.0027}},{"id":"shot-20-18260306-2434476306","idx":20,"score":103,"poi":{"x":-0.0009,"y":-0.0013}},{"id":"shot-21-18270159-2434482159","idx":21,"score":83,"poi":{"x":0.0054,"y":-0.0039}},{"id":"shot-22-18275362-2434487362","idx":22,"score":108,"poi":{"x":-0.0001,"y":0.0003}},{"id":"shot-23-18283476-2434491476","idx":23,"score":98,"poi":{"x":0.0001,"y":-0.0029}},{"id":"shot-24-18335811-2434523811","idx":24,"score":92,"poi":{"x":-0.0031,"y":-0.0031}},{"id":"shot-25-18344837-2434528837","idx":25,"score":96,"poi":{"x":0.0027,"y":-0.0019}},{"id":"shot-26-18353590-2434533590","idx":26,"score":92,"poi":{"x":-0.0017,"y":-0.0039}},{"id":"shot-27-18362828-2434538828","idx":27,"score":60,"poi":{"x":-0.0124,"y":-0.0011}},{"id":"shot-28-18371756-2434543756","idx":28,"score":87,"poi":{"x":-0.0053,"y":-0.0018}},{"id":"shot-29-18381568-2434549568","idx":29,"score":105,"poi":{"x":-0.0005,"y":0.0009}},{"id":"shot-30-18385503-2434553503","idx":30,"score":102,"poi":{"x":0.0006,"y":-0.0016}},{"id":"shot-31-18394790-2434558790","idx":31,"score":86,"poi":{"x":-0.0055,"y":-0.0019}},{"id":"shot-32-18403399-2434563399","idx":32,"score":84,"poi":{"x":-0.0064,"y":-0.0007}},{"id":"shot-33-18412078-2434568078","idx":33,"score":77,"poi":{"x":-0.0076,"y":-0.0026}},{"id":"shot-34-18422082-2434574082","idx":34,"score":83,"poi":{"x":-0.0063,"y":0.0017}},{"id":"shot-35-18494049-2434618049","idx":35,"score":78,"poi":{"x":0.0077,"y":-0.0011}},{"id":"shot-36-18503484-2434623484","idx":36,"score":88,"poi":{"x":-0.0043,"y":-0.0033}},{"id":"shot-37-18512972-2434628972","idx":37,"score":78,"poi":{"x":-0.0049,"y":-0.0063}},{"id":"shot-38-18522385-2434634385","idx":38,"score":63,"poi":{"x":0.0105,"y":-0.0048}},{"id":"shot-39-18531651-2434639651","idx":39,"score":93,"poi":{"x":-0.004,"y":-0.0012}},{"id":"shot-40-18534696-2434642696","idx":40,"score":83,"poi":{"x":0,"y":-0.0067}},{"id":"shot-41-18542802-2434646802","idx":41,"score":78,"poi":{"x":-0.0038,"y":-0.007}},{"id":"shot-42-18550274-2434650274","idx":42,"score":97,"poi":{"x":-0.0013,"y":-0.0028}},{"id":"shot-43-18554152-2434654152","idx":43,"score":99,"poi":{"x":-0.0022,"y":-0.0015}},{"id":"shot-44-18562364-2434658364","idx":44,"score":100,"poi":{"x":-0.0018,"y":-0.0017}},{"id":"shot-45-18571056-2434663056","idx":45,"score":102,"poi":{"x":-0.0005,"y":-0.0017}},{"id":"shot-46-18575934-2434667934","idx":46,"score":100,"poi":{"x":-0.0014,"y":-0.002}},{"id":"shot-47-18584568-2434672568","idx":47,"score":102,"poi":{"x":0.0019,"y":0}},{"id":"shot-48-18592787-2434676787","idx":48,"score":100,"poi":{"x":-0.0023,"y":0.0004}},{"id":"shot-49-19001707-2434681707","idx":49,"score":76,"poi":{"x":0.006,"y":-0.006}},{"id":"shot-50-19010524-2434686524","idx":50,"score":104,"poi":{"x":0.0004,"y":-0.0012}},{"id":"shot-51-19015028-2434691028","idx":51,"score":101,"poi":{"x":0.0018,"y":0.0012}},{"id":"shot-52-19072425-2434724425","idx":52,"score":107,"poi":{"x":-0.0004,"y":-0.0005}},{"id":"shot-53-19080789-2434728789","idx":53,"score":75,"poi":{"x":0.0084,"y":-0.0024}},{"id":"shot-54-19085704-2434733704","idx":54,"score":100,"poi":{"x":-0.0004,"y":-0.0024}},{"id":"shot-55-19094389-2434738389","idx":55,"score":86,"poi":{"x":0.0058,"y":0.0013}},{"id":"shot-56-19102183-2434742183","idx":56,"score":88,"poi":{"x":0.0045,"y":0.0032}},{"id":"shot-57-19110679-2434746679","idx":57,"score":87,"poi":{"x":-0.0054,"y":0.0017}},{"id":"shot-58-19115445-2434751445","idx":58,"score":106,"poi":{"x":0.0008,"y":0.0001}},{"id":"shot-59-19124664-2434756664","idx":59,"score":97,"poi":{"x":0.0025,"y":0.0019}},{"id":"shot-60-19134029-2434762029","idx":60,"score":98,"poi":{"x":0.0029,"y":0.0007}}]`

// 2025-11-06
const shotsJSON string = `[{"id":"shot-1-18181904-2676349904","idx":1,"score":108,"poi":{"x":0.0003,"y":0.0002}},{"id":"shot-2-18185862-2676353862","idx":2,"score":91,"poi":{"x":-0.002,"y":-0.0041}},{"id":"shot-3-18193270-2676357270","idx":3,"score":102,"poi":{"x":0.0012,"y":0.0014}},{"id":"shot-4-18201125-2676361125","idx":4,"score":91,"poi":{"x":-0.0037,"y":0.0028}},{"id":"shot-5-18203916-2676363916","idx":5,"score":90,"poi":{"x":-0.0045,"y":-0.0018}},{"id":"shot-6-18213061-2676369061","idx":6,"score":102,"poi":{"x":0.0007,"y":0.0017}},{"id":"shot-7-18221263-2676373263","idx":7,"score":102,"poi":{"x":-0.0011,"y":-0.0016}},{"id":"shot-8-18225281-2676377281","idx":8,"score":104,"poi":{"x":-0.0004,"y":0.0013}},{"id":"shot-9-18232738-2676380738","idx":9,"score":67,"poi":{"x":-0.0105,"y":0.002}},{"id":"shot-10-18240287-2676384287","idx":10,"score":95,"poi":{"x":-0.0035,"y":0.001}},{"id":"shot-11-18245337-2676389337","idx":11,"score":90,"poi":{"x":0.0048,"y":0.0007}},{"id":"shot-12-18253406-2676393406","idx":12,"score":94,"poi":{"x":-0.0037,"y":-0.0009}},{"id":"shot-13-18261485-2676397485","idx":13,"score":102,"poi":{"x":-0.0018,"y":0.0001}},{"id":"shot-14-18265041-2676401041","idx":14,"score":98,"poi":{"x":-0.0001,"y":-0.0028}},{"id":"shot-15-18272673-2676404673","idx":15,"score":92,"poi":{"x":0.0044,"y":0.0001}},{"id":"shot-16-18281080-2676409080","idx":16,"score":99,"poi":{"x":0.0024,"y":-0.0009}},{"id":"shot-17-18284668-2676412668","idx":17,"score":92,"poi":{"x":-0.0037,"y":0.0022}},{"id":"shot-18-18293243-2676417243","idx":18,"score":81,"poi":{"x":0.0068,"y":-0.0021}},{"id":"shot-19-18301243-2676421243","idx":19,"score":90,"poi":{"x":0.0041,"y":0.0026}},{"id":"shot-20-18305182-2676425182","idx":20,"score":99,"poi":{"x":0.0026,"y":0}},{"id":"shot-21-18333008-2676441008","idx":21,"score":92,"poi":{"x":-0.0036,"y":-0.0024}},{"id":"shot-22-18341236-2676445236","idx":22,"score":100,"poi":{"x":0.0023,"y":0.0008}},{"id":"shot-23-18345691-2676449691","idx":23,"score":107,"poi":{"x":-0.0002,"y":-0.0006}},{"id":"shot-24-18353066-2676453066","idx":24,"score":100,"poi":{"x":0.002,"y":0.0013}},{"id":"shot-25-18362409-2676458409","idx":25,"score":105,"poi":{"x":-0.0009,"y":-0.0007}},{"id":"shot-26-18370909-2676462909","idx":26,"score":90,"poi":{"x":0.0045,"y":0.0019}},{"id":"shot-27-18373826-2676465826","idx":27,"score":91,"poi":{"x":0.0025,"y":0.0038}},{"id":"shot-28-18382313-2676470313","idx":28,"score":86,"poi":{"x":0.0045,"y":0.0037}},{"id":"shot-29-18390549-2676474549","idx":29,"score":87,"poi":{"x":0.0044,"y":-0.0036}},{"id":"shot-30-18400160-2676480160","idx":30,"score":85,"poi":{"x":-0.006,"y":0.0013}},{"id":"shot-31-18405179-2676485179","idx":31,"score":107,"poi":{"x":-0.0003,"y":0.0005}},{"id":"shot-32-18412005-2676488005","idx":32,"score":100,"poi":{"x":-0.0022,"y":0.0009}},{"id":"shot-33-18420350-2676492350","idx":33,"score":93,"poi":{"x":0.004,"y":-0.0005}},{"id":"shot-34-18430646-2676498646","idx":34,"score":91,"poi":{"x":-0.0046,"y":0.0003}},{"id":"shot-35-18434629-2676502629","idx":35,"score":86,"poi":{"x":0.0051,"y":-0.0029}},{"id":"shot-36-18441647-2676505647","idx":36,"score":91,"poi":{"x":0.0023,"y":0.0041}},{"id":"shot-37-18450268-2676510268","idx":37,"score":102,"poi":{"x":0.0016,"y":-0.0011}},{"id":"shot-38-18453677-2676513677","idx":38,"score":88,"poi":{"x":-0.0055,"y":-0.0006}},{"id":"shot-39-18461110-2676517110","idx":39,"score":80,"poi":{"x":0.004,"y":-0.0061}},{"id":"shot-40-18465053-2676521053","idx":40,"score":100,"poi":{"x":0.0023,"y":0.0005}},{"id":"shot-41-18485479-2676533479","idx":41,"score":95,"poi":{"x":0.0023,"y":-0.0028}},{"id":"shot-42-18493390-2676537390","idx":42,"score":103,"poi":{"x":0.0013,"y":-0.0007}},{"id":"shot-43-18500795-2676540795","idx":43,"score":93,"poi":{"x":-0.0041,"y":0.0004}},{"id":"shot-44-18504176-2676544176","idx":44,"score":102,"poi":{"x":0.0018,"y":0.0004}},{"id":"shot-45-18512118-2676548118","idx":45,"score":100,"poi":{"x":0.002,"y":-0.0013}},{"id":"shot-46-18515594-2676551594","idx":46,"score":86,"poi":{"x":0.0056,"y":0.0016}},{"id":"shot-47-18523989-2676555989","idx":47,"score":105,"poi":{"x":0.0002,"y":0.0011}},{"id":"shot-48-18530562-2676558562","idx":48,"score":97,"poi":{"x":0.0025,"y":0.002}},{"id":"shot-49-18533581-2676561581","idx":49,"score":94,"poi":{"x":0.0035,"y":-0.0016}},{"id":"shot-50-18540813-2676564813","idx":50,"score":102,"poi":{"x":-0.0003,"y":0.0018}},{"id":"shot-51-18545459-2676569459","idx":51,"score":92,"poi":{"x":0.0035,"y":-0.0026}},{"id":"shot-52-18553544-2676573544","idx":52,"score":88,"poi":{"x":-0.004,"y":0.0037}},{"id":"shot-53-18561070-2676577070","idx":53,"score":97,"poi":{"x":0.0031,"y":-0.0002}},{"id":"shot-54-18565612-2676581612","idx":54,"score":103,"poi":{"x":0.0016,"y":-0.0005}},{"id":"shot-55-18573291-2676585291","idx":55,"score":95,"poi":{"x":-0.0036,"y":-0.0003}},{"id":"shot-56-18580914-2676588914","idx":56,"score":100,"poi":{"x":-0.0022,"y":-0.0009}},{"id":"shot-57-18585927-2676593927","idx":57,"score":86,"poi":{"x":0.0059,"y":0.0011}},{"id":"shot-58-18592942-2676596942","idx":58,"score":87,"poi":{"x":0.005,"y":-0.0028}},{"id":"shot-59-18595691-2676599691","idx":59,"score":98,"poi":{"x":-0.0027,"y":-0.0009}},{"id":"shot-60-19003160-2676603160","idx":60,"score":98,"poi":{"x":-0.0027,"y":0.0007}}]`
