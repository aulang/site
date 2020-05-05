package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Link struct {
	Title string `json:"title"` // 标题
	Url   string `json:"url"`   // 链接
	Desc  string `json:"desc"`  // 描述
}

type WebConfig struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Title    string             `json:"title"`    // 标题
	Desc     string             `json:"desc"`     // 描述
	Keywords string             `json:"keywords"` // 搜索关键字
	Author   string             `json:"author"`   // 作者
	Website  string             `json:"website"`  // 网站
	Email    string             `json:"email"`    // 邮件
	Github   string             `json:"github"`   // GitHub
	WeChat   string             `json:"wechat"`   // 微信
	Avatar   string             `json:"avatar"`   // 头像
	Since    string             `json:"since"`    // 开始年份
	Links    []Link             `json:"links"`    // 友情链接
}
