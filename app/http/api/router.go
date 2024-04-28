package api

import (
	"gin-api/app/ctx"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type HandlerFunc func(*Ctx) (interface{}, error)

type IRouter interface {
	Use(...gin.HandlerFunc) IRouter
	Group(path string) IRouter

	GET(string, HandlerFunc) IRouter
	POST(string, HandlerFunc) IRouter
	PUT(string, HandlerFunc) IRouter
	DELETE(string, HandlerFunc) IRouter
}

type Router struct {
	gin.IRouter
	ctx *ctx.Ctx
}

func NewRouter(router gin.IRouter, ctx *ctx.Ctx) *Router {
	return &Router{
		IRouter: router,
		ctx:     ctx,
	}
}

func (r *Router) Inject(apis ...IApi) {
	for _, i := range apis {
		i.SetRouter(r)
	}
}

func (r *Router) Use(args ...gin.HandlerFunc) IRouter {
	r.IRouter.Use(args...)
	return r
}

func (r *Router) Group(path string) IRouter {
	return NewRouter(r.IRouter.Group(path), r.ctx)
}

func (r *Router) GET(path string, handler HandlerFunc) IRouter {
	r.IRouter.GET(path, func(ctx *gin.Context) {
		r.handle(ctx, handler)
	})
	return r
}

func (r *Router) POST(path string, handler HandlerFunc) IRouter {
	r.IRouter.POST(path, func(ctx *gin.Context) {
		r.handle(ctx, handler)
	})
	return r
}

func (r *Router) PUT(path string, handler HandlerFunc) IRouter {
	r.IRouter.PUT(path, func(ctx *gin.Context) {
		r.handle(ctx, handler)
	})
	return r
}

func (r *Router) DELETE(path string, handler HandlerFunc) IRouter {
	r.IRouter.DELETE(path, func(ctx *gin.Context) {
		r.handle(ctx, handler)
	})
	return r
}

func (r *Router) handle(ctx *gin.Context, handler HandlerFunc) {
	var resp Response

	if data, err := handler(&Ctx{Gin: ctx, App: r.ctx}); err != nil {
		log.Err(err).Msgf("path: %s", ctx.Request.URL.Path)
		resp = r.error(err)
	} else {
		resp = NewResponse(data)
	}

	ctx.JSON(resp.Status, resp)
}

func (r *Router) error(err error) Response {
	switch e := err.(type) {
	case *Error:
		return NewErrorResponse(e)
	case validator.ValidationErrors:
		return NewErrorResponse(ErrorParameter)
	}

	return NewErrorResponse(ErrorServer)
}
