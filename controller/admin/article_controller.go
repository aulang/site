package admin

import (
	"github.com/kataras/iris/v12"
	"site/entity"
	. "site/model"
	"site/service"
	"time"
)

type ArticleController struct {
	Ctx            iris.Context
	ArticleService service.ArticleService
	CommentService service.CommentService
}

// POST /admin/articles
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
