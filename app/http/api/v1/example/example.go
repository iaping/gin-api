package example

import "gin-api/app/http/api"

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
}
