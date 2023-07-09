package article

import (
	"context"
	"golang_demo/models"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{
		db: db,
	}
}

func (r *ArticleRepository) Find(ctx context.Context) ([]*models.Article, error) {
	var articles []*models.Article
	if err := r.db.WithContext(ctx).Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}
