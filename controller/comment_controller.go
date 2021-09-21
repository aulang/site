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

// GetTopBy /comment/top/{num:int64}
func (c *CommentController) GetTopBy(num int64) Response {
	comments, err := c.CommentService.GetTop(num)

	if err != nil {
		log.Printf("查询最新评论失败，%v", err)
		return FailWithError(err)
	}

	return SuccessWithData(comments)
}

// Post /comment
func (c *CommentController) Post() Response {
	var comment entity.Comment

	if err := c.Ctx.ReadJSON(&comment); err != nil {
		return FailWithCodeAndError(400, err)
	}

	comment.CreationDate = time.Now()

	err := c.CommentService.Save(&comment)

	if err != nil {
		log.Printf("保存评论失败，%v", err)
		return FailWithError(err)
	}

	return SuccessWithData(comment)
}

// PostByReply /comment/{commentId:string}/reply
func (c *CommentController) PostByReply(commentId string) Response {
	var reply entity.Reply

	if err := c.Ctx.ReadJSON(&reply); err != nil {
		return FailWithCodeAndError(400, err)
	}

	reply.CreationDate = time.Now()

	comment, err := c.CommentService.Reply(commentId, &reply)

	if err != nil {
		log.Printf("保存回复失败，%v", err)
		return FailWithError(err)
	}

	return SuccessWithData(comment)
}
