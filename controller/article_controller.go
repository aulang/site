package controller

import (
	"errors"
	"log"

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

// GET /articles/page
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

	var results []interface{}

	for _, data := range page.Datas {
		article, ok := data.(entity.Article)

		if !ok {
			return FailWithError(errors.New("类型转换错误"))
		}

		articleId := article.ID.Hex()

		comments, err := c.CommentService.FindByArticleId(articleId)

		if err != nil {
			log.Printf("查询文章评论失败，%v", err)
		}

		results = append(results, ArticleComment{Article: article, Comments: comments})
	}

	page.Datas = results

	return SuccessWithData(page)
}
