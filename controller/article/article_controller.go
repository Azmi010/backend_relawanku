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

	user, err := articleController.articleServiceInterface.CreateArticle(articleCreated.CreateArticleToModel())
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, response.CreateArticleFromModel(user))
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
