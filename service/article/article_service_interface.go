package article

import "backend_relawanku/model"

type ArticleServiceInterface interface {
	CreateArticle(article model.Article) (model.Article, error)
}