package request

import (
	"backend_relawanku/model"
	"time"

	"gorm.io/gorm"
)

type UpdateArticleRequest struct {
	Title     string    `json:"title" form:"title"`
	Content   string    `json:"content" form:"content"`
	Category  string    `json:"category" form:"category"`
	ImageUrl  string    `json:"image_url" form:"image_url"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

func (updateArticleRequest UpdateArticleRequest) UpdateArticleToModel() model.Article {
	return model.Article{
		Title: updateArticleRequest.Title,
		Content: updateArticleRequest.Content,
		Category: updateArticleRequest.Category,
		ImageUrl: updateArticleRequest.ImageUrl,
		Model: gorm.Model{
			UpdatedAt: updateArticleRequest.UpdatedAt,
		},
	}
}