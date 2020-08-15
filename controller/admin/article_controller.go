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
