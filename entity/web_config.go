package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Link struct {
	Title string `json:"title" bson:"title"`                   // 标题
	URL   string `json:"url" bson:"url"`                       // 链接
	Desc  string `json:"desc,omitempty" bson:"desc,omitempty"` // 描述
}

type WebConfig struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Title        string             `json:"title" bson:"title"`                   // 标题
	Desc         string             `json:"desc,omitempty" bson:"desc,omitempty"` // 描述
	Keywords     string             `json:"keywords" bson:"keywords"`             // 搜索关键字
	Author       string             `json:"author" bson:"author"`                 // 作者
	Website      string             `json:"website" bson:"website"`               // 网站
	Email        string             `json:"email" bson:"email"`                   // 邮件
	Github       string             `json:"github" bson:"github"`                 // GitHub
	WeChat       string             `json:"wechat" bson:"wechat"`                 // 微信
	WeChatQRCode string             `json:"wechatQRCode" bson:"wechatQRCode"`     // 微信二维码
	Avatar       string             `json:"avatar" bson:"avatar"`                 // 头像
	Since        string             `json:"since" bson:"since"`                   // 开始年份
	Links        []Link             `json:"links" bson:"links"`                   // 友情链接
}
