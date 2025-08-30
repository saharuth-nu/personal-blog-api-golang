package repositories

import (
	"blog-api/core/models"

	"gorm.io/gorm"
)

type articleRepositoryDB struct {
	db *gorm.DB
}

func NewArticleRepositoryDB(db *gorm.DB) ArticleRepository {
	return articleRepositoryDB{db: db}
}

func (r articleRepositoryDB) GetAll() ([]Article, error) {
	articles := []Article{}
	tx := r.db.Order("id").Find(&articles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return articles, nil
}

func (r articleRepositoryDB) GetByCreateAt(articleFilter models.ArticleReqFilter) ([]Article, error) {
	articles := []Article{}
	tx := r.db.Order("id").
		Where("created_at BETWEEN ? AND ?", articleFilter.StartPublishingDate, articleFilter.LastPublishingDate).
		Find(&articles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return articles, nil
}

func (r articleRepositoryDB) GetByTag(articleFilter models.ArticleReqFilter) ([]Article, error) {
	articles := []Article{}
	tx := r.db.Order("id").
		Where("tag = ?", articleFilter.Tag).
		Find(&articles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return articles, nil
}

func (r articleRepositoryDB) GetByFilter(articleFilter models.ArticleReqFilter) ([]Article, error) {
	articles := []Article{}
	tx := r.db.Order("id").
		Where("tag = ? AND created_at BETWEEN ? AND ?", articleFilter.Tag, articleFilter.StartPublishingDate, articleFilter.LastPublishingDate).
		Find(&articles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return articles, nil
}

func (r articleRepositoryDB) GetByUID(uid string) (*Article, error) {
	article := Article{ArticleUID: uid}
	tx := r.db.First(&article)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &article, nil
}

func (r articleRepositoryDB) Create(article Article) error {
	tx := r.db.Create(&article)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r articleRepositoryDB) DeleteByUID(uid string) error {
	// soft delete
	tx := r.db.Where("article_uid = ?", uid).Delete(&Article{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r articleRepositoryDB) UpdateByUID(uid string, article models.NewArticleRequest) error {
	// Update with struct
	tx := r.db.Model(Article{}).Where("article_uid = ?", uid).Updates(article)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
