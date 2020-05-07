package controller

import (
	"log"
	. "site/model"
	"site/service"
)

type CategoryController struct {
	categoryService service.CategoryService
}

// GET /categories
func (c *CategoryController) Get() Response {
	categories, err := c.categoryService.GetAll()

	if err != nil {
		log.Printf("查询菜单失败，%v", err)
	}

	return SuccessWithData(categories)
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		categoryService: service.NewCategoryService(),
	}
}
