package models

import "time"

type NewArticleRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type ArticleReqFilter struct {
	Tag                 string    `json:"tag"`
	StartPublishingDate time.Time `json:"start_publishing_date"`
	LastPublishingDate  time.Time `json:"last_publishing_date"`
}
