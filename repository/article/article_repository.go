package article

import (
	"backend_relawanku/model"
	"time"

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

func (articleRepo ArticleRepo) UpdateArticle(articleId uint, article model.Article) (model.Article, error) {
	articleDb := FromModelArticle(article)
	result := articleRepo.db.First(&articleDb, articleId)
	if result.Error != nil {
		return model.Article{}, result.Error
	}

	articleDb.Title = article.Title
	articleDb.Content = article.Content
	articleDb.Category = article.Category
	articleDb.ImageUrl = article.ImageUrl
	articleDb.UpdatedAt = time.Now()

	saveResult := articleRepo.db.Save(&articleDb)
	if saveResult.Error != nil {
		return model.Article{}, saveResult.Error
	}

	return articleDb.ToModelArticle(), nil
}

func (articleRepo ArticleRepo) DeleteArticle(articleId uint) error {
	var articleDb Article
	result := articleRepo.db.First(&articleDb, articleId)
	if result.Error != nil {
		return result.Error
	}

	deleteResult := articleRepo.db.Delete(&articleDb)
	return deleteResult.Error
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

func (articleRepo ArticleRepo) GetTrendingArticles() ([]model.Article, error) {
	var articles []Article
	result := articleRepo.db.Order("view DESC").Find(&articles)
	if result.Error != nil {
		return nil, result.Error
	}

	var modelArticles []model.Article
	for _, a := range articles {
		modelArticles = append(modelArticles, a.ToModelArticle())
	}
	return modelArticles, nil
}

func (articleRepo ArticleRepo) IncrementArticleView(id uint) error {
	return articleRepo.db.Model(&model.Article{}).
		Where("id = ?", id).
		UpdateColumn("View", gorm.Expr("View + 1")).Error
}