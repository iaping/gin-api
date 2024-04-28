package app

import (
	"gin-api/app/config"
	"gin-api/app/ctx"
	"gin-api/app/http"
)

type App struct {
	ctx   *ctx.Ctx
	serve *http.Serve
}

func New(cfg *config.Config) *App {
	return &App{
		ctx: ctx.New(cfg),
	}
}

func (app *App) Run() (err error) {
	if err = app.ctx.Init(); err != nil {
		return
	}

	app.init()

	return app.serve.Run()
}

func (app *App) init() {
	app.serve = http.New(app.ctx)
}
