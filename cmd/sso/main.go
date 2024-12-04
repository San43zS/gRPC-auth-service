package main

import (
	"github.com/op/go-logging"
	"sso/internal/app"
	"sso/internal/config"
)

var log = logging.MustGetLogger("sso")

func main() {
	config := config.MustLoad()

	log.Infof("parsing config: %v", config)

	application := app.New(config.GRPC.Port, config.StoragePath, config.TokenTTL)
	application.Server.MustRun()
}
