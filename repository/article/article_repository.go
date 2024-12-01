package article

import (
	"backend_relawanku/model"

	"gorm.io/gorm"
)

func NewArticleRepository(db *gorm.DB) *ArticleRepo {
	return &ArticleRepo{
		db: db,
	}
}

type ArticleRepo struct {
	db *gorm.DB
}

func (articleRepo ArticleRepo) CreateArticle(article model.Article) (model.Article, error) {
	articleDb := FromModelArticle(article)
	result := articleRepo.db.Create(&articleDb)
	if result.Error != nil {
		return model.Article{}, result.Error
	}
	return articleDb.ToModelArticle(), nil
}