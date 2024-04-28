package middleware

import (
	"fmt"
	"gin-api/app/ctx"

	"github.com/gin-gonic/gin"
)

type Test struct {
	app *ctx.Ctx
}

func NewTest(app *ctx.Ctx) *Test {
	return &Test{
		app: app,
	}
}

func (i *Test) Handle(ctx *gin.Context) {
	fmt.Println("I'm a middleware, debug:", i.app.Config.Debug)
	ctx.Next()
}
