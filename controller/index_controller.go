package controller

import . "github.com/aulang/site/model"

type IndexController struct {
}

// GET /
func (c *IndexController) Get() Response {
	return Success()
}
