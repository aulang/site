package model

type Link struct {
	Title       string //标题
	Url         string //链接
	Description string //描述
}

type BeiAn struct {
	GXBNo  string //工信部备案号，鄂ICP备18028762号
	GXBUrl string //工信部查询地址，http://beian.miit.gov.cn
	GABNo  string //公安部备案号，鄂公网安备 42011102003752号
	GABUrl string //公安部查询地址，http://www.beian.gov.cn/portal/registerSystemInfo?recordcode=42011102003752
}

// 系统配置
type Config struct {
	ID          string
	WebTitle    string // 网站标题
	Keywords    string // 搜索关键字
	Description string // 网站描述
	Since       string // 开始年份
	BeiAn       BeiAn  // 备案号
	Author      string // 网站作者
	Email       string // 邮箱
	GitHub      string // GitHub
	AvatarUrl   string // 头像地址
	Links       []Link //链接
}
