package example

import (
	"gin-api/app/http/api"
)

// @Tags Example
// @Summary 中间件
// @Produce json
// @Router /v1/example/middleware [get]
// @Success 200
func (i *Example) Middleware(ctx *api.Ctx) (interface{}, error) {
	return "I used a middleware", nil
}
