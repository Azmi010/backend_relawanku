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
		ArticleServiceInterface: as,
	}
}

type ArticleController struct {
	ArticleServiceInterface article.ArticleServiceInterface
}

// @Summary      Buat Artikel Baru
// @Description  Membuat artikel baru oleh admin
// @Tags         articles
// @Accept       json
// @Produce      json
// @Param        article  body      request.CreateArticleRequest  true  "Informasi Artikel"
// @Success      201      {object}  map[string]interface{}
// @Router       /api/v1/admin/article [post]
// @Security     BearerAuth
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

	createdArticle, err := articleController.ArticleServiceInterface.CreateArticle(articleCreated.CreateArticleToModel(), file, fileHeader)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, response.CreateArticleFromModel(createdArticle))
}

// @Summary      Update Artikel
// @Description  Memperbarui artikel berdasarkan ID
// @Tags         articles
// @Accept       json
// @Produce      json
// @Param        id       path      uint                   true  "ID Artikel"
// @Param        article  body      request.UpdateArticleRequest  true  "Informasi Artikel yang Diperbarui"
// @Success      200      {object}  map[string]interface{}
// @Router       /api/v1/admin/article/{id} [put]
// @Security     BearerAuth
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

	file, fileHeader, err := c.Request().FormFile("image_url")
	if err != nil {
		return base.ErrorResponse(c, errors.New("failed to read image file"))
	}
	defer file.Close()

	updatedArticle, err := articleController.ArticleServiceInterface.UpdateArticle(uint(id), updateRequest.UpdateArticleToModel(), file, fileHeader)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, response.UpdateArticleFromModel(updatedArticle))
}

// @Summary      Hapus Artikel
// @Description  Menghapus artikel berdasarkan ID
// @Tags         articles
// @Param        id  path      uint  true  "ID Artikel"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/admin/article/{id} [delete]
// @Security     BearerAuth
func (articleController ArticleController) DeleteArticleController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return base.ErrorResponse(c, errors.New("invalid article ID"))
	}

	err = articleController.ArticleServiceInterface.DeleteArticle(uint(id))
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, "Article deleted successfully")
}

// @Summary      Dapatkan Semua Artikel
// @Description  Mengambil daftar semua artikel
// @Tags         articles
// @Produce      json
// @Success      200  {array}   response.ArticleResponse
// @Router       /api/v1/admin/articles [get]
// @Security     BearerAuth
func (articleController ArticleController) GetAllArticlesController(c echo.Context) error {
	articles, err := articleController.ArticleServiceInterface.GetAllArticles()
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, articles)
}

// @Summary      Dapatkan Artikel Sesuai Kategori
// @Description  Mengambil daftar semua artikel sesuai kategori
// @Tags         articles
// @Param 		 category path string true "Category Article"
// @Produce      json
// @Success      200  {array}   response.ArticleResponse
// @Router       /api/v1/admin/articles/{category} [get]
// @Security     BearerAuth
func (articleController ArticleController) GetArticlesByCategoryController(c echo.Context) error {
	category := c.QueryParam("category")
	articles, err := articleController.ArticleServiceInterface.GetArticlesByCategory(category)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, articles)
}

// @Summary      Dapatkan Artikel Sesuai ID
// @Description  Mengambil daftar semua artikel sesuai ID
// @Tags         articles
// @Param 		 id path uint true "Category ID"
// @Sec
// @Produce      json
// @Success      200  {array}   response.ArticleResponse
// @Router       /api/v1/user/articles/{id} [get]
// @Security     BearerAuth
func (articleController ArticleController) GetArticleByIDController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return base.ErrorResponse(c, errors.New("invalid article ID"))
	}
	article, err := articleController.ArticleServiceInterface.GetArticleByID(uint(id))
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, response.ArticleFromModel(article))
}

// @Summary      Dapatkan Artikel Trending
// @Description  Mengambil daftar semua artikel urut sesuai trending
// @Tags         articles
// @Param 		 category path string true "Trending Article"
// @Produce      json
// @Success      200  {array}   response.ArticleResponse
// @Router       /api/v1/user/article-trending [get]
// @Security     BearerAuth
func (articleController ArticleController) GetTrendingArticlesController(c echo.Context) error {
	articles, err := articleController.ArticleServiceInterface.GetTrendingArticles()
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, articles)
}