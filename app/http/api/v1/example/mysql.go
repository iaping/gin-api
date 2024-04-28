package example

import (
	"context"
	"gin-api/app/http/api"
	"gin-api/app/model"

	"github.com/uptrace/bun"
)

// @Tags Example
// @Summary mysql示例
// @Produce json
// @Router /v1/example/mysql [get]
// @Success 200
func (i *Example) Mysql(ctx *api.Ctx) (interface{}, error) {
	type timeZoneName struct {
		model.Model
		bun.BaseModel `bun:"table:time_zone_name,alias:tzn"`
		TimeZoneId    int    `json:"time_zone_id"`
		Name          string `json:"name"`
	}

	var data []timeZoneName
	q := ctx.App.Db.NewSelect().Model(&data)
	q.Limit(15)
	q.Order("Time_zone_id ASC")
	err := q.Scan(context.Background())

	return data, err
}
