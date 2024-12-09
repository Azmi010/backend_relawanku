package dashboard

import (
	"backend_relawanku/controller/article"
	"backend_relawanku/controller/base"
	"backend_relawanku/controller/program"

	"github.com/labstack/echo/v4"
)

type DashboardController struct {
	articleController *article.ArticleController
	programController *program.ProgramController
}

func NewDashboardController(articleCtrl *article.ArticleController, programCtrl *program.ProgramController) *DashboardController {
	return &DashboardController{
		articleController: articleCtrl,
		programController: programCtrl,
	}
}

func (ctrl *DashboardController) GetDashboardData(c echo.Context) error {
	articles, err := ctrl.articleController.ArticleServiceInterface.GetAllArticles()  
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	programs, err := ctrl.programController.Service.GetAllPrograms() 
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, map[string]interface{}{
		"articles": articles,
		"programs": programs,
	})
}
