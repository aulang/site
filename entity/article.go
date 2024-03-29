package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Article 文章
type Article struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Private       bool               `json:"private" bson:"private"`               // 是否私有
	Title         string             `json:"title" bson:"title"`                   // 标题
	SubTitle      string             `json:"subTitle" bson:"subTitle"`             // 副标题
	Summary       string             `json:"summary" bson:"summary"`               // 简述
	Content       string             `json:"content" bson:"content"`               // html
	Source        string             `json:"source" bson:"source"`                 // 原内容
	Renew         time.Time          `json:"renew" bson:"renew"`                   // 更新时间
	CategoryID    string             `json:"categoryId" bson:"categoryId"`         // 类别ID
	CategoryName  string             `json:"categoryName" bson:"categoryName"`     // 类别名
	CreationDate  time.Time          `json:"creationDate" bson:"creationDate"`     // 创建日期
	CommentsCount int                `json:"commentsCount" bson:"commentsCount"`   // 评论计数
	Tags          []string           `bson:"tags,omitempty" json:"tags,omitempty"` // 标签
}
