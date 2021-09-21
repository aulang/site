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

// Post /admin/article
func (c *ArticleController) Post() Response {
	var article entity.Article

	if err := c.Ctx.ReadJSON(&article); err != nil {
		return FailWithError(err)
	}

	now := time.Now()
	if article.ID.IsZero() {
		article.CreationDate, article.Renew = now, now
	} else {
		db, err := c.ArticleService.GetByID(article.ID.Hex())
		if err != nil {
			return FailWithError(err)
		}

		db.Renew = now

		db.Title = article.Title
		db.SubTitle = article.SubTitle
		db.Summary = article.Summary
		db.Content = article.Content
		db.Source = article.Source
		db.CategoryID = article.CategoryID
		db.CategoryName = article.CategoryName

		article = db
	}

	err := c.ArticleService.Save(&article)

	if err != nil {
		return FailWithError(err)
	}

	return SuccessWithData(article)
}

//DeleteBy DELETE /admin/article/{id}
func (c *ArticleController) DeleteBy(id string) Response {
	err := c.ArticleService.Delete(id)
	if err != nil {
		return FailWithError(err)
	} else {
		return Success()
	}
}

//GetPage GET /admin/article/page?page=1&size=20&keyword=keyword&category=category
func (c *ArticleController) GetPage() Response {
	page := c.Ctx.URLParamInt64Default("page", 1)
	size := c.Ctx.URLParamInt64Default("size", 20)

	if page < 1 {
		page = 1
	}

	if size < 1 {
		size = 20
	}

	keyword := c.Ctx.URLParam("keyword")
	category := c.Ctx.URLParam("category")

	result, err := c.ArticleService.GetByPage(page, size, keyword, category)
	if err != nil {
		return FailWithError(err)
	}

	return SuccessWithData(result)
}
