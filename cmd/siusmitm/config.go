package main

import (
	"context"

	"github.com/diwise/service-chassis/pkg/infrastructure/servicerunner"
	"github.com/istyf/siusmitm/pkg/application"
)

type FlagType int
type FlagMap map[FlagType]string

const (
	listenAddress FlagType = iota
	webAssets
	webPort
	siusClientPort
	siusServerPort
	loadResults
)

type AppConfig struct {
	app *application.App

	SiusClientPort string
	SiusServerPort string
	WebAssetsPath  string
	WebPort        string

	ResultFilePath string

	cancelContext context.CancelFunc
}

var onstarting = servicerunner.OnStarting[AppConfig]
var onrunning = servicerunner.OnRunning[AppConfig]
var onshutdown = servicerunner.OnShutdown[AppConfig]
var webserver = servicerunner.WithHTTPServeMux[AppConfig]
var muxinit = servicerunner.OnMuxInit[AppConfig]
var listen = servicerunner.WithListenAddr[AppConfig]
var port = servicerunner.WithPort[AppConfig]
