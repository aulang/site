package controller

import (
	. "github.com/aulang/site/model"
	"github.com/aulang/site/service"
	"log"
)

type WebConfigController struct {
	ConfigService service.WebConfigService
}

// Get /config
func (c *WebConfigController) Get() Response {
	config, err := c.ConfigService.Get()

	if err != nil {
		log.Printf("查询配置信息失败，%v", err)
		return FailWithError(err)
	}

	return SuccessWithData(config.Sort())
}
