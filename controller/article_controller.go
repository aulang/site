package controller

import (
	"github.com/kataras/iris/v12"
	"log"
	. "site/model"
	"site/service"
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

	return SuccessWithData(article)
}

// GET /articles/top3
func (c *ArticleController) GetTop3() Response {
	articles, err := c.ArticleService.GetTop3()

	if err != nil {
		log.Printf("查询最新文章失败，%v", err)
	}

	return SuccessWithData(articles)
}

// GET /articles/pages
func (c *ArticleController) GetPage() Response {
	pageNo := c.Ctx.URLParamIntDefault("page", 1)
	pageSize := c.Ctx.URLParamIntDefault("size", 1)

	articles, err := c.ArticleService.Page(int64(pageNo), int64(pageSize))
	if err != nil {
		return FailWithError(err)
	}

	if articles == nil {
		return Success()
	}

	var results []ArticleComment

	for _, article := range articles {
		articleId := article.ID.Hex()

		comments, err := c.CommentService.FindByArticleId(articleId)

		if err != nil {
			log.Printf("查询文章评论失败，%v", err)
		}

		if comments != nil {
			article.CommentsCount += len(comments)
		}

		results = append(results, ArticleComment{Article: article, Comments: comments})
	}

	return SuccessWithData(results)
}
