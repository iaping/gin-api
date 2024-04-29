package job

import "gin-api/app/ctx"

type IJob interface {
	Spec() string
	Run(*ctx.Ctx)
}
