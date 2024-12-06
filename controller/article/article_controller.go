package article

import (
	"backend_relawanku/controller/article/request"
	"backend_relawanku/controller/article/response"
	"backend_relawanku/controller/base"
	"backend_relawanku/service/article"
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewArticleController(as article.ArticleServiceInterface) *ArticleController {
	return &ArticleController{
		articleServiceInterface: as,
	}
}

type ArticleController struct {
	articleServiceInterface article.ArticleServiceInterface
}

func (articleController ArticleController) CreateArticleController(c echo.Context) error {
	articleCreated := request.CreateArticleRequest{}
	if err := c.Bind(&articleCreated); err != nil {
		return base.ErrorResponse(c, err)
	}

	file, fileHeader, err := c.Request().FormFile("image_url")
	if err != nil {
		return base.ErrorResponse(c, errors.New("failed to read image file"))
	}
	defer file.Close()

	createdArticle, err := articleController.articleServiceInterface.CreateArticle(articleCreated.CreateArticleToModel(), file, fileHeader)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, response.CreateArticleFromModel(createdArticle))
}

func (articleController ArticleController) UpdateArticleController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return base.ErrorResponse(c, errors.New("invalid article ID"))
	}

	updateRequest := request.UpdateArticleRequest{}
	if err := c.Bind(&updateRequest); err != nil {
		return base.ErrorResponse(c, err)
	}

	updatedArticle, err := articleController.articleServiceInterface.UpdateArticle(uint(id), updateRequest.UpdateArticleToModel())
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, response.UpdateArticleFromModel(updatedArticle))
}

func (articleController ArticleController) DeleteArticleController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return base.ErrorResponse(c, errors.New("invalid article ID"))
	}

	err = articleController.articleServiceInterface.DeleteArticle(uint(id))
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, "Article deleted successfully")
}

func (articleController ArticleController) GetAllArticlesController(c echo.Context) error {
	articles, err := articleController.articleServiceInterface.GetAllArticles()
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, articles)
}

func (articleController ArticleController) GetArticlesByCategoryController(c echo.Context) error {
	category := c.QueryParam("category")
	articles, err := articleController.articleServiceInterface.GetArticlesByCategory(category)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, articles)
}

func (articleController ArticleController) GetArticleByIDController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return base.ErrorResponse(c, errors.New("invalid article ID"))
	}
	article, err := articleController.articleServiceInterface.GetArticleByID(uint(id))
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, article)
}
