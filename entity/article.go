package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// 文章
type Article struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Title         string             `json:"title"`         // 标题
	SubTitle      string             `json:"subTitle"`      // 副标题
	Summary       string             `json:"summary"`       // 总结
	Content       string             `json:"content"`       // 内容
	Renew         time.Time          `json:"renew"`         // 更新时间
	Category      Category           `json:"category"`      // 类别
	CreationDate  time.Time          `json:"creationDate"`  // 创建日期
	CommentsCount int                `json:"commentsCount"` // 评论计数
}
