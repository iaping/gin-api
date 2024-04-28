package api

import (
	"gin-api/app/ctx"

	"github.com/gin-gonic/gin"
)

type Ctx struct {
	Gin *gin.Context
	App *ctx.Ctx
}
