package article

import (
	"context"
	"golang_demo/models"
)

type ArticleRepository interface {
	Find(ctx context.Context) ([]*models.Article, error)
}
