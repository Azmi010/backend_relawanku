package request

import (
	"backend_relawanku/model"

	"gorm.io/gorm"
)

type CreateArticleRequest struct {
	Title     string         `json:"title" form:"title"`
	Content   string         `json:"content" form:"content"`
	Category  string         `json:"category" form:"category"`
	ImageUrl  string         `json:"image_url" form:"image_url"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}

func (createArticleRequest CreateArticleRequest) CreateArticleToModel() model.Article {
	return model.Article{
		Title:    createArticleRequest.Title,
		Content:  createArticleRequest.Content,
		Category: createArticleRequest.Category,
		ImageUrl: createArticleRequest.ImageUrl,
	}
}
