package example

import (
	"gin-api/app/http/api"
	"gin-api/app/http/middleware"
)

type Example struct {
	api.Api
}

func (i *Example) SetRouter(r api.IRouter) {
	rg := r.Group("/example")

	rg.GET("/helloworld", i.Helloworld)
	rg.GET("/error", i.Error)
	rg.GET("/error/custom", i.ErrorCustom)
	rg.GET("/redis", i.Redis)
	rg.GET("/mysql", i.Mysql)

	rg.Use(middleware.NewTest(r.App()).Handle)
	rg.GET("/middleware", i.Middleware)
}
