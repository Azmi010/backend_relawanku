package article

import (
	"backend_relawanku/controller/article/request"
	"backend_relawanku/controller/article/response"
	"backend_relawanku/controller/base"
	"backend_relawanku/service/article"

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