package controller

import . "site/model"

type IndexController struct {
}

// GET /
func (c *IndexController) Get() Response {
	return Success()
}

func NewIndexController() *IndexController {
	return &IndexController{}
}
