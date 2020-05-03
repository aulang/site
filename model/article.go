package model

import (
	"site/entity"
)

type Article struct {
	ID            string
	Title         string
	SubTitle      string
	Summary       string
	Content       string
	Renew         string
	CategoryID    string
	CategoryName  string
	CommentsCount int
}

func NewArticle(article *entity.Article) Article {
	return Article{
		ID:            article.ID.Hex(),
		Title:         article.Title,
		SubTitle:      article.SubTitle,
		Summary:       article.Summary,
		Content:       article.Summary,
		Renew:         article.Renew.String(),
		CategoryID:    article.Category.ID.Hex(),
		CategoryName:  article.Category.Name,
		CommentsCount: article.CommentsCount,
	}
}
