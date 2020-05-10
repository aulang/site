package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// 回复
type Reply struct {
	Mail         string    `json:"mail"`         // 回复人邮件
	Name         string    `json:"name"`         // 回复人
	Content      string    `json:"content"`      // 内容
	CreationDate time.Time `json:"creationDate"` // 创建日期
}

// 评论
type Comment struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Mail         string             `json:"mail"`         // 评论人邮件
	Name         string             `json:"name"`         // 评论人
	ArticleId    string             `json:"articleId"`    // 文章ID
	Content      string             `json:"content"`      // 内容
	CreationDate time.Time          `json:"creationDate"` // 创建日期
	Replies      []Reply            `json:"replies"`      // 回复
}
