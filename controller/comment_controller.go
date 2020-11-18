package controller

import (
	"log"
	"time"

	"github.com/aulang/site/entity"
	. "github.com/aulang/site/model"
	"github.com/aulang/site/service"
	"github.com/kataras/iris/v12"
)

type CommentController struct {
	Ctx            iris.Context
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

// POST /comment
func (c *CommentController) Post() Response {
	var comment entity.Comment

	if err := c.Ctx.ReadJSON(&comment); err != nil {
		return FailWithError(err)
	}

	comment.CreationDate = time.Now()

	err := c.CommentService.Save(&comment)

	if err != nil {
		return FailWithError(err)
	}

	return SuccessWithData(comment)
}

// POST /comment/{commentId:string}/reply
func (c *CommentController) PostByReply(commentId string) Response {
	var reply entity.Reply

	if err := c.Ctx.ReadJSON(&reply); err != nil {
		return FailWithError(err)
	}

	reply.CreationDate = time.Now()

	comment, err := c.CommentService.Reply(commentId, &reply)

	if err != nil {
		return FailWithError(err)
	}

	return SuccessWithData(comment)
}
