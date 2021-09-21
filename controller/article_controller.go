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

// GetBy GET /article/{id:string}
func (c *ArticleController) GetBy(id string) Response {
	article, err := c.ArticleService.GetByID(id)

	if err != nil {
		log.Printf("查询文章信息失败，%v", err)
		return FailWithError(err)
	}

	articleId := article.ID.Hex()
	comments, err := c.CommentService.FindByArticleId(articleId)

	if err != nil {
		log.Printf("查询文章评论失败，%v", err)
		return FailWithError(err)
	}

	ac := ArticleComment{Article: article, Comments: comments}

	return SuccessWithData(ac)
}

// GetTopBy GET /article/top/{num:int64}
func (c *ArticleController) GetTopBy(num int64) Response {
	articles, err := c.ArticleService.GetTop(num)

	if err != nil {
		log.Printf("查询最新文章失败，%v", err)
		return FailWithError(err)
	}

	return SuccessWithData(articles)
}

// GetPage /article/page?page=1&size=20&keyword=keyword&category=category
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
		log.Printf("分页查询文章失败，%v", err)
		return FailWithError(err)
	}

	return SuccessWithData(result)
}
