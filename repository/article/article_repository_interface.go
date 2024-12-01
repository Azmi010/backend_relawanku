package article

import "backend_relawanku/model"

type ArticleRepository interface {
	CreateArticle(article model.Article) (model.Article, error)
}