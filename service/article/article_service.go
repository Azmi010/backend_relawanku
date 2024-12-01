package article

import (
	"backend_relawanku/model"
	articleRepo "backend_relawanku/repository/article"
	"errors"
)

func NewArticleService(ar articleRepo.ArticleRepository) *ArticleService {
	return  &ArticleService{
		articleRepoInterface: ar,
	}
}

type ArticleService struct {
	articleRepoInterface articleRepo.ArticleRepository
}

func (articleService ArticleService) CreateArticle(article model.Article) (model.Article, error) {
	createdArticle, err := articleService.articleRepoInterface.CreateArticle(article)
	if err != nil {
		return model.Article{}, errors.New("failed to create article")
	}

	return createdArticle, nil
}