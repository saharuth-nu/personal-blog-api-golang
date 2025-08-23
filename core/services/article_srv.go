package services

import (
	"blog-api/core/models"
	"blog-api/core/repositories"
	"errors"

	"github.com/google/uuid"
)

type articleService struct {
	articleRepo repositories.ArticleRepository
}

func NewArticleService(articleRepo repositories.ArticleRepository) ArticleService {
	return articleService{articleRepo: articleRepo}
}

func (s articleService) GetArticles() ([]ArticleResponse, error) {

	articles, err := s.articleRepo.GetAll()
	if err != nil {
		return nil, err
	}

	articleResponses := []ArticleResponse{}
	for _, article := range articles {
		articleResponse := ArticleResponse{
			ArticleUID:     article.ArticleUID,
			Title:          article.Title,
			Content:        article.Content,
			Tag:            article.Tag,
			PublishingDate: article.CreatedAt,
		}
		articleResponses = append(articleResponses, articleResponse)
	}

	return articleResponses, nil
}

func (s articleService) GetArticleByUID(uid string) (*ArticleResponse, error) {

	if err := uuid.Validate(uid); err != nil {
		return nil, err
	}

	article, err := s.articleRepo.GetByUID(uid)
	if err != nil {
		return nil, err
	}

	articleResponse := ArticleResponse{
		ArticleUID:     article.ArticleUID,
		Title:          article.Title,
		Content:        article.Content,
		Tag:            article.Tag,
		PublishingDate: article.CreatedAt,
	}

	return &articleResponse, nil
}

func (s articleService) CreateArticle(request models.NewArticleRequest) error {

	if request.Content == "" {
		return errors.New("content is empty")
	}

	article := repositories.Article{
		ArticleUID: uuid.New().String(),
		Title:      request.Title,
		Content:    request.Content,
		Tag:        request.Tag,
	}

	if err := s.articleRepo.Create(article); err != nil {
		return err
	}

	return nil
}

func (s articleService) DeleteArticleByUID(uid string) error {

	if err := uuid.Validate(uid); err != nil {
		return err
	}

	_, err := s.articleRepo.GetByUID(uid)
	if err != nil {
		return err
	}

	if err := s.articleRepo.DeleteByUID(uid); err != nil {
		return err
	}

	return nil
}
