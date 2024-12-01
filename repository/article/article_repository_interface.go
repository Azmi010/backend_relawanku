package article

import "backend_relawanku/model"

type ArticleRepository interface {
	CreateArticle(article model.Article) (model.Article, error)
	GetAllArticles() ([]model.Article, error)
	GetArticlesByCategory(category string) ([]model.Article, error)
	GetArticleByID(id uint) (model.Article, error)
}
