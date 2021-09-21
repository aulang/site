package model

import "github.com/aulang/site/entity"

// ArticleComment 文章和评论
type ArticleComment struct {
	entity.Article                  // 文章
	Comments       []entity.Comment `json:"comments,omitempty"` // 评论
}
