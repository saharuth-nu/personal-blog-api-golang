package services

import (
	"blog-api/core/models"
	"blog-api/core/repositories"
	"errors"
	"log"

	"github.com/google/uuid"
)

type articleService struct {
	articleRepo repositories.ArticleRepository
}

func NewArticleService(articleRepo repositories.ArticleRepository) ArticleService {
	return articleService{articleRepo: articleRepo}
}

func (s articleService) GetArticlesByFilter(articleFilter models.ArticleReqFilter) ([]ArticleResponse, error) {

	filterTag := false
	filterDate := false

	log.Println(articleFilter.LastPublishingDate)
	log.Println(articleFilter.StartPublishingDate)

	if !articleFilter.LastPublishingDate.IsZero() && !articleFilter.StartPublishingDate.IsZero() {
		filterDate = true
	}

	if articleFilter.Tag != "" {
		filterTag = true
	}

	log.Println(filterDate)

	var articles []repositories.Article
	var err error

	if filterTag && filterDate {
		articles, err = s.articleRepo.GetByFilter(articleFilter)
	} else if filterTag {
		articles, err = s.articleRepo.GetByTag(articleFilter)
	} else if filterDate {
		articles, err = s.articleRepo.GetByCreateAt(articleFilter)
	} else {
		articles, err = s.articleRepo.GetAll()
	}

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

	if request.Content == "" || request.Title == "" {
		return errors.New("title or content is empty")
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

func (s articleService) UpdateArticleByUID(uid string, request models.NewArticleRequest) error {

	if err := uuid.Validate(uid); err != nil {
		return err
	}

	_, err := s.articleRepo.GetByUID(uid)
	if err != nil {
		return err
	}

	if request.Content == "" || request.Title == "" {
		return errors.New("title or content is empty")
	}

	if err := s.articleRepo.UpdateByUID(uid, request); err != nil {
		return err
	}

	return nil
}
