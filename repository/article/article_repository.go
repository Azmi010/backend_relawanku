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

func (articleRepo ArticleRepo) GetAllArticles() ([]model.Article, error) {
	var articles []Article
	result := articleRepo.db.Find(&articles)
	if result.Error != nil {
		return nil, result.Error
	}

	var modelArticles []model.Article
	for _, a := range articles {
		modelArticles = append(modelArticles, a.ToModelArticle())
	}
	return modelArticles, nil
}

func (articleRepo ArticleRepo) GetArticlesByCategory(category string) ([]model.Article, error) {
	var articles []Article
	result := articleRepo.db.Where("category = ?", category).Find(&articles)
	if result.Error != nil {
		return nil, result.Error
	}

	var modelArticles []model.Article
	for _, a := range articles {
		modelArticles = append(modelArticles, a.ToModelArticle())
	}
	return modelArticles, nil
}

func (articleRepo ArticleRepo) GetArticleByID(id uint) (model.Article, error) {
	var article Article
	result := articleRepo.db.First(&article, id)
	if result.Error != nil {
		return model.Article{}, result.Error
	}
	return article.ToModelArticle(), nil
}
