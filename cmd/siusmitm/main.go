package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/diwise/frontend-toolkit/pkg/middleware/csp"
	"github.com/diwise/service-chassis/pkg/infrastructure/env"
	"github.com/diwise/service-chassis/pkg/infrastructure/o11y/logging"
	"github.com/diwise/service-chassis/pkg/infrastructure/servicerunner"

	"github.com/istyf/siusmitm/pkg/api"
	"github.com/istyf/siusmitm/pkg/mitm"
)

func DefaultFlags() FlagMap {
	webAssetPath := ""
	if cwd, err := os.Getwd(); err == nil {
		webAssetPath = filepath.ToSlash(filepath.Join(cwd, "assets"))
	}

	return FlagMap{
		listenAddress:  "", // listen on all ipv4 and ipv6 interfaces
		webAssets:      webAssetPath,
		webPort:        "8088", //
		siusClientPort: "4300", // default sius comm service port
		siusServerPort: "4200", // default sius comm server port
		loadResults:    "",
	}
}

func main() {
	ctx, flags := parseExternalConfig(context.Background(), DefaultFlags())

	logger := slog.Default()
	ctx = logging.NewContextWithLogger(ctx, logger)

	cfg, err := newConfig(ctx, flags)
	exitIf(err, logger, "failed to create application config")

	ctx, cfg.cancelContext = context.WithCancel(ctx)

	runner, err := initialize(ctx, flags, cfg)
	exitIf(err, logger, "failed to initialize service")

	err = runner.Run(ctx, servicerunner.WithWorker(func(ctx_ context.Context, cfg *AppConfig) error {

		if cfg.ResultFilePath != "" {
			f, err := os.Open(cfg.ResultFilePath)
			if err != nil {
				logger.Error("failed to open results file", "err", err.Error())
			} else {
				shotsJSON, err := io.ReadAll(f)
				f.Close()

				if err == nil {
					var shotsToAdd []mitm.Shot
					json.Unmarshal([]byte(shotsJSON), &shotsToAdd)

					for len(shotsToAdd) > 0 {
						time.Sleep(time.Duration(500) * time.Millisecond)
						s := shotsToAdd[0]
						shotsToAdd = shotsToAdd[1:]

						if jsonBytes, err := json.Marshal(s); err == nil {
							_, err := http.Post(
								"http://localhost:"+cfg.WebPort+"/api/shots",
								"application/json",
								bytes.NewBuffer(jsonBytes),
							)
							if err != nil {
								logger.Error("failed to post shot data", "err", err.Error())
							}
						}
					}
				}
			}
		}

		logger.Info("listening for connections on", "port", cfg.SiusClientPort)

		listener, err := net.Listen("tcp", ":"+cfg.SiusClientPort)
		if err != nil {
			return err
		}

		for context.Cause(ctx_) != context.Canceled {
			inconn, err := listener.Accept()
			if err != nil {
				return err
			}

			logger.Info("accepted connection on port", "port", cfg.SiusClientPort)

			tcpAddr, err := net.ResolveTCPAddr("tcp", ":"+cfg.SiusServerPort)
			if err != nil {
				return fmt.Errorf("resolve tcp addr failed: %s", err.Error())
			}

			outconn, err := net.DialTCP("tcp", nil, tcpAddr)
			if err != nil {
				return fmt.Errorf("dial failed: %s", err.Error())
			}

			logger.Info("connected successfully to port", "port", cfg.SiusServerPort)

			SetupPipes(ctx_, inconn, outconn)
		}

		return nil
	}))
	exitIf(err, logger, "service runner failed")

	logger.Info("shutting down")
}

func SetupPipes(ctx context.Context, inconn net.Conn, outconn net.Conn) {
	go mitm.Pipe(ctx, bufio.NewReader(inconn), outconn)
	go mitm.Pipe(ctx, bufio.NewReader(outconn), inconn)
}

func parseExternalConfig(ctx context.Context, flags FlagMap) (context.Context, FlagMap) {

	// Allow environment variables to override certain defaults
	envOrDef := env.GetVariableOrDefault
	flags[webPort] = envOrDef(ctx, "SERVICE_PORT", flags[webPort])
	flags[siusClientPort] = envOrDef(ctx, "PORT", flags[siusClientPort])
	flags[siusServerPort] = envOrDef(ctx, "SIUS_PORT", flags[siusServerPort])

	apply := func(f FlagType) func(string) error {
		return func(value string) error {
			flags[f] = value
			return nil
		}
	}

	// Allow command line arguments to override defaults and environment variables
	flag.Func("siusport", "sius comm service port number", apply(siusServerPort))
	flag.Func("port", "sius server port", apply(siusClientPort))
	flag.Func("web-port", "web service port number", apply(webPort))
	flag.Func("web-assets", "web assets path", apply(webAssets))
	flag.Func("load", "load results from file", apply(loadResults))
	flag.Parse()

	return ctx, flags
}

func newConfig(_ context.Context, flags FlagMap) (*AppConfig, error) {
	cfg := &AppConfig{
		SiusClientPort: flags[siusClientPort],
		SiusServerPort: flags[siusServerPort],
		ResultFilePath: flags[loadResults],
	}
	return cfg, nil
}

func initialize(ctx context.Context, flags FlagMap, cfg *AppConfig) (servicerunner.Runner[AppConfig], error) {

	_, runner := servicerunner.New(ctx, *cfg,
		webserver("public", listen(flags[listenAddress]), port(flags[webPort]),
			muxinit(func(ctx context.Context, identifier string, port string, svcCfg *AppConfig, handler *http.ServeMux) (err error) {
				svcCfg.WebPort = port

				middlewares := append(
					make([]func(http.Handler) http.Handler, 0, 10),
					csp.NewContentSecurityPolicy(csp.StrictDynamic()),
				)

				assetLoader, err := api.NewAssetLoader(ctx, flags[webAssets])
				if err != nil {
					return err
				}

				err = api.RegisterHandlers(ctx, handler, middlewares, assetLoader, svcCfg.app)
				if err != nil {
					return fmt.Errorf("failed to create new api handler: %s", err.Error())
				}

				return nil
			}),
		),
		onstarting(func(ctx context.Context, svcCfg *AppConfig) (err error) {
			return nil
		}),
		onrunning(func(ctx context.Context, svcCfg *AppConfig) error {
			logging.GetFromContext(ctx).Info("service is running and waiting for connections")
			return nil
		}),
		onshutdown(func(ctx context.Context, svcCfg *AppConfig) error {
			if svcCfg.cancelContext != nil {
				svcCfg.cancelContext()
				svcCfg.cancelContext = nil
			}

			return nil
		}),
	)

	return runner, nil
}

func exitIf(err error, logger *slog.Logger, msg string, args ...any) {
	if err != nil {
		logger.With(args...).Error(msg, "err", err.Error())
		os.Exit(1)
	}
}
