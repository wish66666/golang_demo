package article_test

import (
	"context"
	"golang_demo/controllers/article"
	"golang_demo/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ArticleRepositoryMock struct {
	mock.Mock
}

func (r *ArticleRepositoryMock) Find(ctx context.Context) ([]*models.Article, error) {
	result := r.Called()
	return result.Get(0).([]*models.Article), result.Error(1)
}

type ArticleTestSuite struct {
	suite.Suite
	echo *echo.Echo
	mock *ArticleRepositoryMock
}

func (s *ArticleTestSuite) SetupTest() {
	s.echo = echo.New()
	s.mock = new(ArticleRepositoryMock)
	controller := article.NewArticleController(s.mock)
	s.echo.GET("/articles", controller.GetAll)
}

func (s *ArticleTestSuite) TestGetAll() {
	articles := []*models.Article{
		{
			ID:        1,
			Title:     "hoge",
			Content:   "fuga",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	s.mock.On("Find").Return(articles, nil)

	req := httptest.NewRequest(http.MethodGet, "/articles", nil)
	rec := httptest.NewRecorder()

	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)
}

func TestArticleTestSuite(t *testing.T) {
	suite.Run(t, new(ArticleTestSuite))
}
