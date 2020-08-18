package model

type Page struct {
	PageNo     int64         `json:"pageNo"`
	PageSize   int64         `json:"pageSize"`
	Datas      []interface{} `json:"datas"`
	TotalPages int64         `json:"totalPages"`
}
