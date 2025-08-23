package services

import (
	"blog-api/core/models"
	"time"
)

type ArticleResponse struct {
	ArticleUID     string    `json:"article_uid"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Tag            string    `json:"tag"`
	PublishingDate time.Time `json:"publishing_date"`
}

type ArticleService interface {
	GetArticles() ([]ArticleResponse, error)
	GetArticleByUID(string) (*ArticleResponse, error)
	CreateArticle(models.NewArticleRequest) error
	DeleteArticleByUID(string) error
}
