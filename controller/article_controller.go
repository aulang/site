package controller

import (
	"log"

	. "github.com/aulang/site/model"
	"github.com/aulang/site/service"
	"github.com/kataras/iris/v12"
)

type ArticleController struct {
	Ctx            iris.Context
	ArticleService service.ArticleService
	CommentService service.CommentService
}

// GET /articles/{id:string}
func (c *ArticleController) GetBy(id string) Response {
	article, err := c.ArticleService.GetByID(id)

	if err != nil {
		return Fail(-1, "记录不存在")
	}

	articleId := article.ID.Hex()
	comments, err := c.CommentService.FindByArticleId(articleId)

	if err != nil {
		log.Printf("查询文章评论失败，%v", err)
	}

	ac := ArticleComment{Article: article, Comments: comments}

	return SuccessWithData(ac)
}

// GET /articles/top3
func (c *ArticleController) GetTop3() Response {
	articles, err := c.ArticleService.GetTop3()

	if err != nil {
		log.Printf("查询最新文章失败，%v", err)
	}

	return SuccessWithData(articles)
}

// GET /articles/page
func (c *ArticleController) GetPage() Response {
	var defaultValue int64 = 1

	pageNo := c.Ctx.URLParamInt64Default("page", defaultValue)
	pageSize := c.Ctx.URLParamInt64Default("size", defaultValue)

	if pageNo < 1 {
		pageNo = 1
	}

	if pageSize < 1 {
		pageSize = 1
	}

	keyword := c.Ctx.URLParam("keyword")
	category := c.Ctx.URLParam("category")

	page, err := c.ArticleService.GetByPage(pageNo, pageSize, keyword, category)
	if err != nil {
		return FailWithError(err)
	}

	return SuccessWithData(page)
}
