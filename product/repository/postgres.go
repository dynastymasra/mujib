package repository

import (
	"context"

	"github.com/dynastymasra/mujib/domain"
	"github.com/jinzhu/gorm"
)

const (
	TableName = "products"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Save(ctx context.Context, product domain.Product) error {
	return r.db.Omit("created_at").Table(TableName).Save(&product).Error
}

func (r *ProductRepository) FindByID(ctx context.Context, id string) (*domain.Product, error) {
	var (
		result domain.Product
		query  = domain.Product{
			ID: id,
		}
	)

	err := r.db.Table(TableName).Where(query).First(&result).Error

	return &result, err
}

func (r *ProductRepository) Fetch(ctx context.Context) ([]domain.Product, error) {
	var result []domain.Product

	err := r.db.Table(TableName).Find(&result).Error

	return result, err
}

func (r *ProductRepository) Update(ctx context.Context, product domain.Product) error {
	return r.db.Omit("created_at").Table(TableName).Update(&product).Error
}

func (r *ProductRepository) Delete(ctx context.Context, product domain.Product) error {
	return r.db.Table(TableName).Delete(product).Error
}
