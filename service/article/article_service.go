package article

import (
	"backend_relawanku/helper"
	"backend_relawanku/model"
	articleRepo "backend_relawanku/repository/article"
	"errors"
	"mime/multipart"
)

func NewArticleService(ar articleRepo.ArticleRepository) *ArticleService {
	return  &ArticleService{
		articleRepoInterface: ar,
	}
}

type ArticleService struct {
	articleRepoInterface articleRepo.ArticleRepository
}

func (articleService ArticleService) CreateArticle(article model.Article, file multipart.File, fileHeader *multipart.FileHeader) (model.Article, error) {
	imageURL, err := helper.UploadImageToFirebase("my-chatapp-01.appspot.com", "articles", fileHeader.Filename, file)
	if err != nil {
		return model.Article{}, errors.New("failed to upload image to Firebase")
	}

	article.ImageUrl = imageURL

	createdArticle, err := articleService.articleRepoInterface.CreateArticle(article)
	if err != nil {
		return model.Article{}, errors.New("failed to create article")
	}

	return createdArticle, nil
}

func (articleService ArticleService) UpdateArticle(id uint, article model.Article, file multipart.File, fileHeader *multipart.FileHeader) (model.Article, error) {
	if file != nil && fileHeader != nil {
		imageURL, err := helper.UploadImageToFirebase("my-chatapp-01.appspot.com", "articles", fileHeader.Filename, file)
		if err != nil {
			return model.Article{}, errors.New("failed to upload image to Firebase")
		}
		article.ImageUrl = imageURL
	}

	updated, err := articleService.articleRepoInterface.UpdateArticle(id, article)
	if err != nil {
		return model.Article{}, errors.New("failed to update article")
	}
	return updated, nil
}

func (articleService ArticleService) DeleteArticle(articleId uint) error {
	err := articleService.articleRepoInterface.DeleteArticle(articleId)
	if err != nil {
		return errors.New("failed to delete article")
	}
	return nil
}

func (articleService ArticleService) GetAllArticles() ([]model.Article, error) {
	return articleService.articleRepoInterface.GetAllArticles()
}

func (articleService ArticleService) GetArticlesByCategory(category string) ([]model.Article, error) {
	return articleService.articleRepoInterface.GetArticlesByCategory(category)
}

func (articleService ArticleService) GetArticleByID(id uint) (model.Article, error) {
	if err := articleService.articleRepoInterface.IncrementArticleView(id); err != nil {
		return model.Article{}, errors.New("failed to increment view count")
	}
	
	article, err := articleService.articleRepoInterface.GetArticleByID(id)
	if err != nil {
		return model.Article{}, errors.New("article not found")
	}

	return article, nil
}

func (articleService ArticleService) GetTrendingArticles() ([]model.Article, error) {
	return articleService.articleRepoInterface.GetTrendingArticles()
}