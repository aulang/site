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

// GET /login
func (c *IndexController) GetLogin() {
	c.Ctx.Redirect(c.OAuth.AuthorizeUrl(), iris.StatusFound)
}

// GET /token
func (c *IndexController) GetToken() Response {
	accessToken := c.OAuth.Token(c.Ctx)

	if accessToken == nil {
		return Fail(500, "获取令牌失败")
	}

	user, err := c.OAuth.User(accessToken)
	if err != nil {
		return FailWithError(err)
	}

	// TODO JWT

	return SuccessWithData(user)
}
