package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 文章类别
type Category struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name"`         // 类别名称
	Count        int32              `json:"count"`        // 类别文章总数
	CreationDate time.Time          `json:"creationDate"` // 评论
}
