package article

import (
	"golang_demo/controllers/interfaces/repositories/article"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ArticleController struct {
	repository article.ArticleRepository
}

func NewArticleController(repository article.ArticleRepository) *ArticleController {
	return &ArticleController{
		repository: repository,
	}
}

func (c *ArticleController) GetAll(ctx echo.Context) error {
	articles, err := c.repository.Find(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, articles)
}
