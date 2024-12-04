package response

import (
	"backend_relawanku/model"
	"time"
)

type UpdateArticleResponse struct {
	ID        uint      `json:"id" form:"id"`
	Title     string    `json:"title" form:"title"`
	Content   string    `json:"content" form:"content"`
	Category  string    `json:"category" form:"category"`
	ImageUrl  string    `json:"image_url" form:"image_url"`
	UpdatedAt time.Time `json:"created_at" form:"created_at"`
}

func UpdateArticleFromModel(article model.Article) UpdateArticleResponse {
	return UpdateArticleResponse{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Category:  article.Category,
		ImageUrl:  article.ImageUrl,
		UpdatedAt: article.UpdatedAt,
	}
}
