package model

type Page struct {
	Size    int64         `json:"size"`    // 每页条数
	Pages   int64         `json:"pages"`   // 总页数
	Total   int64         `json:"total"`   // 总条数
	Current int64         `json:"current"` // 当前页
	Records []interface{} `json:"records"` // 分页记录
}

func NewPage(current, size, total int64, Records []interface{}) *Page {
	var pages int64 = 0
	if size != 0 {
		// 总页数 = (总记录数 + 每页数据大小 - 1) / 每页数据大小
		pages = (total + size - 1) / size
	}

	return &Page{
		Size:    size,
		Pages:   pages,
		Total:   total,
		Current: current,
		Records: Records,
	}
}
