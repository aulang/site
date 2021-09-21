package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Reply 回复
type Reply struct {
	Mail         string    `json:"mail" bson:"mail"`                 // 回复人邮件
	Name         string    `json:"name" bson:"name"`                 // 回复人
	Content      string    `json:"content" bson:"content"`           // 内容
	CreationDate time.Time `json:"creationDate" bson:"creationDate"` // 创建日期
}

// Comment 评论
type Comment struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Mail         string             `json:"mail" bson:"mail"`                           // 评论人邮件
	Name         string             `json:"name" bson:"name"`                           // 评论人
	ArticleID    string             `json:"articleId" bson:"articleId"`                 // 文章ID
	Content      string             `json:"content" bson:"content"`                     // 内容
	CreationDate time.Time          `json:"creationDate" bson:"creationDate"`           // 创建日期
	Replies      []Reply            `json:"replies,omitempty" bson:"replies,omitempty"` // 回复
}
