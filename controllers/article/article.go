package article

import (
	"golang_demo/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ArticleController struct {
	db *gorm.DB
}

func NewArticleController(db *gorm.DB) *ArticleController {
	return &ArticleController{
		db: db,
	}
}

func (c *ArticleController) GetAll(ctx echo.Context) error {
	var articles []models.Article
	if err := c.db.Find(&articles).Error; err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, articles)
}
