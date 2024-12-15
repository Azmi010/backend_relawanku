package article

import (
	"backend_relawanku/model"
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID        uint           `gorm:"primaryKey"`
	Title     string         `gorm:"not null" json:"title" form:"title"`
	SubTitle  string         `gorm:"not null" json:"sub_title" form:"sub_title"`
	Content   string         `gorm:"not null" json:"content" form:"content"`
	Category  string         `gorm:"not null" json:"category" form:"category"`
	View      int            `gorm:"not null" json:"view" form:"view"`
	ImageUrl  string         `gorm:"not null" json:"image_url" form:"image_url"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func FromModelArticle(article model.Article) Article {
	return Article{
		ID:        article.ID,
		Title:     article.Title,
		SubTitle:  article.SubTitle,
		Content:   article.Content,
		Category:  article.Category,
		View:      article.View,
		ImageUrl:  article.ImageUrl,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		DeletedAt: article.DeletedAt,
	}
}

func (article Article) ToModelArticle() model.Article {
	return model.Article{
		Model: gorm.Model{
			ID:        article.ID,
			CreatedAt: article.CreatedAt,
			UpdatedAt: article.UpdatedAt,
			DeletedAt: article.DeletedAt,
		},
		Title:    article.Title,
		SubTitle: article.SubTitle,
		Content:  article.Content,
		Category: article.Category,
		View:     article.View,
		ImageUrl: article.ImageUrl,
	}
}
