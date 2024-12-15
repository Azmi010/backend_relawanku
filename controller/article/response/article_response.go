package response

import (
	"time"
)

type ArticleResponse struct {
	ID        uint      `json:"id" example:"1"`
	Title     string    `json:"title" example:"Article Title"`
	Content   string    `json:"content" example:"Article Content"`
	Category  string    `json:"category" example:"Technology"`
	View      int       `json:"view" example:"100"`
	ImageURL  string    `json:"image_url" example:"https://example.com/image.jpg"`
	CreatedAt time.Time `json:"created_at" example:"2024-12-09T09:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-12-09T09:00:00Z"`
	DeletedAt time.Time `json:"deleted_at" example:"2024-12-09T09:00:00Z"`
}
