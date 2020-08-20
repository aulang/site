package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

// 菜单
type Menu struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`                     // 名称
	URL   string             `json:"url" bson:"url"`                       // 链接
	Desc  string             `json:"desc,omitempty" bson:"desc,omitempty"` // 描述
	Order int                `json:"order" bson:"order"`                   // 排序
}
