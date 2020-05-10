package controller

import (
	"log"
	. "site/model"
	"site/service"
)

type MenuController struct {
	MenuService service.MenuService
}

// GET /menus
func (c *MenuController) Get() Response {
	menus, err := c.MenuService.GetAll()

	if err != nil {
		log.Printf("查询菜单失败，%v", err)
	}

	return SuccessWithData(menus)
}
