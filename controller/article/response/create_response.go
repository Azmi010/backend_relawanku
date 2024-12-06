package response

import (
	"backend_relawanku/model"
	"time"
)

type CreateArticleResponse struct {
	ID        uint      `json:"id" form:"id"`
	Title     string    `json:"title" form:"title"`
	Content   string    `json:"content" form:"content"`
	Category  string    `json:"category" form:"category"`
	ImageUrl  string    `json:"image_url" form:"image_url"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

func CreateArticleFromModel(article model.Article) CreateArticleResponse {
	return CreateArticleResponse{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Category:  article.Category,
		ImageUrl:  article.ImageUrl,
		CreatedAt: article.CreatedAt,
	}
}