package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/istyf/siusmitm/pkg/application"
	"github.com/istyf/siusmitm/pkg/components"
	"github.com/istyf/siusmitm/pkg/mitm"
	"github.com/istyf/siusmitm/pkg/smcontext"

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

func RegisterHandlers(ctx context.Context, mux *http.ServeMux, middleware []func(http.Handler) http.Handler, assetLoader AssetLoader, app *application.App) error {

	version := uuid.NewString()

	r := router.New(mux)

	assets.RegisterEndpoints(ctx, assetLoader, assets.WithMux(mux),
		assets.WithImmutableExpiry(48*time.Hour),
		assets.WithRedirect("/favicon.ico", "/icons/favicon.ico", http.StatusFound),
	)

	shots := make([]mitm.Shot, 0, 100)

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

		ctx := smcontext.SetNameAndClass(r.Context(), "Kalle Kula", "okänd")

		urlValues := r.URL.Query()

		componentURL := "/components/shootlog"
		componentURLParams := []string{}

		if class := urlValues.Get("class"); class != "" {
			if class, err := url.QueryUnescape(class); err == nil {
				componentURLParams = append(componentURLParams, "class="+url.QueryEscape(class))
				ctx = smcontext.SetNameAndClass(ctx, smcontext.Name(ctx), class)
			}
		}

		if name := urlValues.Get("name"); name != "" {
			if name, err := url.QueryUnescape(name); err == nil {
				componentURLParams = append(componentURLParams, "name="+url.QueryEscape(name))
				ctx = smcontext.SetNameAndClass(ctx, name, smcontext.Class(ctx))
			}
		}

		isStanding := true
		if sival := urlValues.Get("si"); sival != "" && sival != "0" && sival != "false" {
			componentURLParams = append(componentURLParams, "si=1")
			isStanding = false
		}

		if len(componentURLParams) > 0 {
			componentURL = componentURL + "?" + strings.Join(componentURLParams, "&")
		}

		mu.Lock()
		defer mu.Unlock()

		shootingLog := components.ShootingLog(version, assetLoader.Load, shots, isStanding, componentURL)
		shootingLog.Render(ctx, w)
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

			ctx := smcontext.SetNameAndClass(r.Context(), "okänd", "okänd")

			urlValues := r.URL.Query()

			if class := urlValues.Get("class"); class != "" {
				if class, err := url.QueryUnescape(class); err == nil {
					ctx = smcontext.SetNameAndClass(ctx, smcontext.Name(ctx), class)
				}
			}

			if name := urlValues.Get("name"); name != "" {
				if name, err := url.QueryUnescape(name); err == nil {
					ctx = smcontext.SetNameAndClass(ctx, name, smcontext.Class(ctx))
				}
			}

			isStanding := true
			if si := urlValues.Get("si"); si == "1" {
				isStanding = false
			}

			mu.Lock()
			defer mu.Unlock()

			card := components.ShootingLogComponent(shots, isStanding)
			card.Render(ctx, w)
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
