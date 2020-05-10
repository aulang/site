package model

import "site/entity"

// 文章和评论
type ArticleComment struct {
	entity.Article
	Comments []entity.Comment `json:"comments"` // 评论
}
