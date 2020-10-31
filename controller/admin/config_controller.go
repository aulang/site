package admin

import (
	. "github.com/aulang/site/model"
	"github.com/aulang/site/service"
	"github.com/kataras/iris/v12"
)

type ConfigController struct {
	Ctx           iris.Context
	MenuService   service.MenuService
	ConfigService service.WebConfigService
}

// POST /admin/config
func (c *ConfigController) Post() Response {
	var config ConfigRequest

	if err := c.Ctx.ReadJSON(&config); err != nil {
		return FailWithError(err)
	}

	webConfig := config.Config
	menus := config.Menus

	for _, menu := range menus {
		err := c.MenuService.Save(&menu)
		if err != nil {
			return FailWithError(err)
		}
	}

	err := c.ConfigService.Save(&webConfig)
	if err != nil {
		return FailWithError(err)
	}

	return Success()
}
