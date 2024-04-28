package http

import (
	"gin-api/app/ctx"
	"gin-api/app/http/api"
	v1 "gin-api/app/http/api/v1"
	"strings"

	_ "gin-api/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Serve struct {
	ctx *ctx.Ctx
	s   *gin.Engine
}

func New(ctx *ctx.Ctx) *Serve {
	if !ctx.Config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	return &Serve{
		ctx: ctx,
		s:   gin.Default(),
	}
}

func (s *Serve) Run() (err error) {
	s.init()
	return s.s.Run(s.ctx.Config.Http.Addr)
}

func (s *Serve) init() {
	s.debug()
	s.cors()
	s.router()
}

func (s *Serve) debug() {
	if !s.ctx.Config.Debug {
		return
	}

	s.s.GET("/_doc/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (s *Serve) cors() {
	cfg := cors.Config{
		AllowOrigins:     strings.Split(s.ctx.Config.Http.Cors.AllowOrigins, ","),
		AllowMethods:     strings.Split(s.ctx.Config.Http.Cors.AllowMethods, ","),
		AllowHeaders:     strings.Split(s.ctx.Config.Http.Cors.AllowHeaders, ","),
		AllowCredentials: s.ctx.Config.Http.Cors.AllowCredentials,
		ExposeHeaders:    strings.Split(s.ctx.Config.Http.Cors.ExposeHeaders, ","),
	}
	s.s.Use(cors.New(cfg))
}

func (s *Serve) router() {
	// your api
	r := api.NewRouter(s.s.Group("/v1"), s.ctx)
	r.Inject([]api.IApi{
		v1.Example,
	}...)
}
