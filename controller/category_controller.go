package controller

import (
	"log"

	. "github.com/aulang/site/model"
	"github.com/aulang/site/service"
)

type CategoryController struct {
	CategoryService service.CategoryService
}

// GET /categories
func (c *CategoryController) Get() Response {
	categories, err := c.CategoryService.GetAll()

	if err != nil {
		log.Printf("查询文章类别失败，%v", err)
	}

	return SuccessWithData(categories)
}
