package model

type Link struct {
	Title string //标题
	Url   string //链接
	Desc  string //描述
}

// 系统配置
type WebConfig struct {
	Title    string // 标题
	Desc     string // 描述
	Keywords string // 搜索关键字
	Author   string // 作者
	Email    string // 邮箱
	GitHub   string // GitHub
	WeChat   string // 微信
	Avatar   string // 头像
	Since    string // 开始年份
	Links    []Link // 链接
}
