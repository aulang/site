package controller

import (
	"github.com/aulang/site/middleware/oauth"
	. "github.com/aulang/site/model"
	"github.com/kataras/iris/v12"
)

type IndexController struct {
	Ctx   iris.Context
	OAuth *oauth.OAuth
}

// GET /
func (c *IndexController) Get() Response {
	return Success()
}
