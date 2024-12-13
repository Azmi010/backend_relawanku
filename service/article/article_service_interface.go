package article

import (
	"backend_relawanku/model"
	"mime/multipart"
)

type ArticleServiceInterface interface {
	CreateArticle(article model.Article, file multipart.File, fileHeader *multipart.FileHeader) (model.Article, error)
	UpdateArticle(articleId uint, article model.Article, file multipart.File, fileHeader *multipart.FileHeader) (model.Article, error)
	DeleteArticle(articleId uint) error
	GetAllArticles() ([]model.Article, error)
	GetArticlesByCategory(category string) ([]model.Article, error)
	GetArticleByID(id uint) (model.Article, error)
}
