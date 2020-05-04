package controller

import (
	"log"
	. "site/model"
	"site/service"
)

type MenuController struct {
	menuService service.MenuService
}

// GET /menus
func (c *MenuController) Get() Response {
	menus, err := c.menuService.GetAll()

	if err != nil {
		log.Printf("查询菜单失败，%v", err)
	}

	return SuccessWithData(menus)
}

func NewMenuController() *MenuController {
	return &MenuController{
		menuService: service.NewMenuService(),
	}
}
