package main

import (
	"gin-api/app"
	"gin-api/app/config"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out: os.Stderr,
	})
}

// @title gin-api
// @version v1
// @description gin-api is a basic api framework developed based on gin
func main() {
	cfg, err := config.Load("./config.yaml")
	if err != nil {
		log.Panic().Err(err).Msg("Config")
	}

	if err = app.New(cfg).Run(); err != nil {
		log.Panic().Err(err).Msg("App")
	}
}
