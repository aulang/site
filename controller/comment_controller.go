package controller

import (
	"log"
	. "site/model"
	"site/service"
)

type CommentController struct {
	commentService service.CommentService
}

// GET /comment/top3
func (c *CommentController) GetTop3() Response {
	articles, err := c.commentService.GetTop3()

	if err != nil {
		log.Printf("查询最进评论失败，%v", err)
	}

	return SuccessWithData(articles)
}

func NewCommentController() *CommentController {
	return &CommentController{
		commentService: service.NewCommentService(),
	}
}
