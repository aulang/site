package controller

import (
	"log"
	. "site/model"
	"site/service"
)

type ArticleController struct {
	articleService service.ArticleService
}

// GET /articles/{id:string}
func (c *ArticleController) GetBy(id string) Response {
	article, err := c.articleService.GetByID(id)

	if err != nil {
		return Fail(-1, "记录不存在")
	}

	return SuccessWithData(article)
}

// GET /articles/top3
func (c *ArticleController) GetTop3() Response {
	articles, err := c.articleService.GetTop3()

	if err != nil {
		log.Printf("查询最新文章失败，%v", err)
	}

	return SuccessWithData(articles)
}

func NewArticleController() *ArticleController {
	return &ArticleController{
		articleService: service.NewArticleService(),
	}
}
