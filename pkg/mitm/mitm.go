package mitm

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Service interface {
	Start(context.Context)
	Stop(context.Context)
}

type impl struct{}

func NewService(ctx context.Context, port string) (Service, error) {
	return &impl{}, nil
}

func (m *impl) Start(context.Context) {}

func (m *impl) Stop(context.Context) {}

type rcData struct {
	Cmd string `xml:"mCmd"`
	Prm string `xml:"mPrm"`
}

func Process(xmlData io.Reader) error {
	decoder := xml.NewDecoder(xmlData)

	for {
		rcd := &rcData{}
		if err := decoder.Decode(rcd); err != nil {
			if !errors.Is(err, io.EOF) {
				return err
			}
			return nil
		}

		if strings.Compare(rcd.Cmd, "Binary") == 0 {
			s, err := DecodeShot(rcd.Prm)
			if err != nil {
				if err == ErrNotAShot {
					continue
				}
				return err
			}

			b, _ := json.Marshal(s)

			fmt.Println("shot:", string(b))

			resp, err := http.Post("http://127.0.0.1:8088/api/shots", "application/json", bytes.NewBuffer(b))
			if err != nil {
				fmt.Println("request failed", err.Error())
				continue
			}

			if resp.StatusCode != 200 {
				fmt.Println("failed to foward shot to web server", resp.StatusCode)
			}
		}
	}
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Shot struct {
	ID string `json:"id"`

	Idx           int64 `json:"idx"`
	ScoreInTenths int64 `json:"score"`
	PointOfImpact Point `json:"poi"`
}

func (s *Shot) Distance() float64 {
	return math.Sqrt((s.PointOfImpact.X * s.PointOfImpact.X) + (s.PointOfImpact.Y * s.PointOfImpact.Y))
}

func (s *Shot) X() float64 {
	return s.PointOfImpact.X
}

func (s *Shot) Y() float64 {
	return s.PointOfImpact.Y
}

func (s *Shot) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

func (s *Shot) TimeOfImpact() time.Time {
	parts := strings.Split(s.ID, "-")

	if len(parts) == 4 && len(parts[2]) == 8 {
		ts := parts[2]
		ts = "2025-06-06T" + ts[0:2] + ":" + ts[2:4] + ":" + ts[4:6] + "." + ts[6:] + "+02:00"

		t, err := time.Parse(time.RFC3339Nano, ts)
		if err != nil {
			fmt.Println("error:", err.Error())
		} else {
			return t
		}
	}

	return time.Time{}
}

var ErrNotAShot = errors.New("invalid shot data")

func DecodeShot(shot string) (Shot, error) {
	parts := strings.Split(shot, ";")
	if len(parts) != 24 || parts[0] != "_SHOT" {
		return Shot{}, ErrNotAShot
	}

	score, err := strconv.ParseInt(parts[10], 10, 64)
	if err != nil {
		return Shot{}, fmt.Errorf("failed to parse shot score from %s: %s", parts[10], err.Error())
	}

	index, err := strconv.ParseInt(parts[13], 10, 64)
	if err != nil {
		return Shot{}, fmt.Errorf("failed to parse shot index from %s: %s", parts[13], err.Error())
	}

	xpoi, err := strconv.ParseFloat(parts[14], 64)
	if err != nil {
		return Shot{}, fmt.Errorf("failed to parse x coordinate of impact from %s: %s", parts[14], err.Error())
	}

	ypoi, err := strconv.ParseFloat(parts[15], 64)
	if err != nil {
		return Shot{}, fmt.Errorf("failed to parse y coordinate of impact from %s: %s", parts[15], err.Error())
	}

	timestamp := strings.ReplaceAll(strings.ReplaceAll(parts[6], ":", ""), ".", "")

	s := Shot{
		ID:            fmt.Sprintf("shot-%s-%s-%s", parts[13], timestamp, parts[20]),
		Idx:           index,
		ScoreInTenths: score,
		PointOfImpact: Point{
			X: xpoi,
			Y: ypoi,
		},
	}

	return s, nil
}

func Pipe(ctx context.Context, in io.Reader, out io.Writer) {

	ch := make(chan byte)

	go func() {
		defer close(ch)

		buf := make([]byte, 32)

		for {
			n, err := in.Read(buf)
			if err != nil {
				return
			}

			for idx := range n {
				ch <- buf[idx]
			}
		}
	}()

	outBuffer := bytes.NewBuffer(make([]byte, 0, 1024))

	for {
		select {
		case b, stillOpen := <-ch:
			{
				if !stillOpen {
					return
				}

				outBuffer.WriteByte(b)

				if b == '>' {
					buf := outBuffer.Bytes()
					buflen := outBuffer.Len()

					endOfFirstTag := slices.Index(buf, '>')
					if buflen >= (2*(endOfFirstTag+1))+1 {
						var isCompleteCommand bool = true

						// Check if the last tag is an endtag with the same length as the first tag
						if buf[buflen-endOfFirstTag-2] == '<' && buf[buflen-endOfFirstTag-1] == '/' {
							// ... if it is, check if the tag is the same
							for chkidx := range endOfFirstTag - 1 {
								if buf[chkidx+1] != buf[buflen-endOfFirstTag+chkidx] {
									isCompleteCommand = false
									break
								}
							}

							if isCompleteCommand {
								if err := Process(bytes.NewBuffer(outBuffer.Bytes())); err != nil {
									fmt.Println("got error", err.Error())
								}

								outBuffer.WriteTo(out)
								outBuffer.Reset()
							}
						}
					}
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
