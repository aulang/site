package controller

import (
	"log"

	. "github.com/aulang/site/model"
	"github.com/aulang/site/service"
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
