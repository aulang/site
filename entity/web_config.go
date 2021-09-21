package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sort"
)

type Link struct {
	Title string `json:"title" bson:"title"`                   // 标题
	URL   string `json:"url" bson:"url"`                       // 链接
	Order int    `json:"order" bson:"order"`                   // 排序
	Desc  string `json:"desc,omitempty" bson:"desc,omitempty"` // 描述
}
type BeiAn struct {
	GxbNo  string `json:"gxbNo" bson:"gxbNo"`   // 工信部备案号
	GxbUrl string `json:"gxbUrl" bson:"gxbUrl"` // 工信部备案链接
	GabNo  string `json:"gabNo" bson:"gabNo"`   // 公安部备案号
	GabUrl string `json:"gabUrl" bson:"gabUrl"` // 公安部备案链接
}

type LinkSlice []Link

func (ls LinkSlice) Len() int           { return len(ls) }
func (ls LinkSlice) Swap(i, j int)      { ls[i], ls[j] = ls[j], ls[i] }
func (ls LinkSlice) Less(i, j int) bool { return ls[i].Order < ls[j].Order }

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
	Menus        LinkSlice          `json:"menus" bson:"menus"`                   // 导航菜单
	Links        LinkSlice          `json:"links" bson:"links"`                   // 友情链接
	BeiAn        BeiAn              `json:"beiAn" bson:"beiAn"`                   // 备案信息
}

func (ws WebConfig) Sort() WebConfig {
	sort.Sort(ws.Links)
	sort.Sort(ws.Menus)
	return ws
}
