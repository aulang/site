package admin

import (
	"time"

	"github.com/aulang/site/entity"
	. "github.com/aulang/site/model"
	"github.com/aulang/site/service"
	"github.com/kataras/iris/v12"
)

type ArticleController struct {
	Ctx            iris.Context
	ArticleService service.ArticleService
	CommentService service.CommentService
}

// POST /admin/article
func (c *ArticleController) Post() Response {
	var article entity.Article

	if err := c.Ctx.ReadJSON(&article); err != nil {
		return FailWithError(err)
	}

	now := time.Now()
	if article.ID.IsZero() {
		article.CreationDate, article.Renew = now, now
	} else {
		article.Renew = now
	}

	err := c.ArticleService.Save(&article)

	if err != nil {
		return FailWithError(err)
	}

	return SuccessWithData(article)
}

// DELETE /admin/article/{id}
func (c *ArticleController) DeleteBy(id string) Response {
	err := c.ArticleService.Delete(id)
	if err != nil {
		return FailWithError(err)
	} else {
		return Success()
	}
}

// GET /admin/article/page
func (c *ArticleController) GetPage() Response {
	pageNo := c.Ctx.URLParamIntDefault("page", 1)
	pageSize := c.Ctx.URLParamIntDefault("size", 1)

	if pageNo < 1 {
		pageNo = 1
	}

	if pageSize < 1 {
		pageSize = 1
	}

	keyword := c.Ctx.URLParam("keyword")
	category := c.Ctx.URLParam("category")

	page, err := c.ArticleService.Page(int64(pageNo), int64(pageSize), keyword, category)
	if err != nil {
		return FailWithError(err)
	}

	return SuccessWithData(page)
}
