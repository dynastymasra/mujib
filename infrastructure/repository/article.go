package repository

import (
	"context"

	"github.com/dynastymasra/mujib/domain"
	"github.com/jinzhu/gorm"
)

const (
	ArticleTableName = "articles"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (a *ArticleRepository) Save(ctx context.Context, article domain.Article) error {
	return a.db.Omit("created_at").Table(ArticleTableName).Save(&article).Error
}

func (a *ArticleRepository) FindByID(ctx context.Context, id string) (*domain.Article, error) {
	var (
		result *domain.Article
		query  = domain.Article{
			ID: id,
		}
	)

	err := a.db.Table(ArticleTableName).Where(query).First(result).Error

	return result, err
}

func (a *ArticleRepository) FindAll(ctx context.Context) ([]*domain.Article, error) {
	var result []*domain.Article

	err := a.db.Table(ArticleTableName).Find(&result).Error

	return result, err
}
