package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

// 菜单
type Menu struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string             `json:"name"` // 名称
	Url   string             `json:"url"`  // 链接
	Desc  string             `json:"desc"` // 描述
	Order int                `json:order`  // 排序
}
