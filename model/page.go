package model

type Page struct {
	PageNo     int64         `json:"pageNo"`
	PageSize   int64         `json:"pageSize"`
	Content    []interface{} `json:"content"`
	TotalPages int64         `json:"totalPages"`
}
