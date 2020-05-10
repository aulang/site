package controller

import (
	"log"
	. "site/model"
	"site/service"
)

type WebConfigController struct {
	ConfigService service.WebConfigService
}

// GET /config
func (c *WebConfigController) Get() Response {
	config, err := c.ConfigService.Get()

	if err != nil {
		log.Printf("查询配置失败，%v", err)
	}

	return SuccessWithData(config)
}
