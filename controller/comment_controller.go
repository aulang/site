package controller

import (
	"log"
	. "site/model"
	"site/service"
)

type CommentController struct {
	CommentService service.CommentService
}

// GET /comment/top3
func (c *CommentController) GetTop3() Response {
	comments, err := c.CommentService.GetTop3()

	if err != nil {
		log.Printf("查询最进评论失败，%v", err)
	}

	return SuccessWithData(comments)
}
