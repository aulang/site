package admin

import (
	"github.com/aulang/site/entity"
	. "github.com/aulang/site/model"
	"github.com/aulang/site/service"
	"github.com/kataras/iris/v12"
)

type ConfigController struct {
	Ctx           iris.Context
	ConfigService service.WebConfigService
}

// POST /admin/config
func (c *ConfigController) Post() Response {
	var config entity.WebConfig

	if err := c.Ctx.ReadJSON(&config); err != nil {
		return FailWithError(err)
	}

	err := c.ConfigService.Save(&config)
	if err != nil {
		return FailWithError(err)
	}

	return Success()
}
