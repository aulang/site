package controller

import (
	. "github.com/aulang/site/model"
)

type IndexController struct {
}

// Get /
func (c *IndexController) Get() Response {
	return Success()
}
