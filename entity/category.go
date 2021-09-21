package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Category 文章类别
type Category struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`                 // 类别名称
	Count        int                `json:"count" bson:"count"`               // 类别文章总数
	Order        int                `json:"order" bson:"order"`               // 排序
	CreationDate time.Time          `json:"creationDate" bson:"creationDate"` // 创建时间
}
