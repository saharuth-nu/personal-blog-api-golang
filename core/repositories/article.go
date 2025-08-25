package repositories

import (
	"blog-api/core/models"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ArticleUID string `gorm:"primaryKey;type:string"`
	Title      string `gorm:"column:title"`
	Content    string `gorm:"column:content"`
	Tag        string `gorm:"column:tag"`
}

type ArticleRepository interface {
	GetAll() ([]Article, error)
	GetByUID(string) (*Article, error)
	Create(Article) error
	DeleteByUID(string) error
	UpdateByUID(string, models.NewArticleRequest) error
}
