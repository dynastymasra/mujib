package test

import (
	"context"

	"github.com/dynastymasra/mujib/domain"
	"github.com/stretchr/testify/mock"
)

type MockArticleRepository struct {
	mock.Mock
}

func (m *MockArticleRepository) Save(ctx context.Context, article domain.Article) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockArticleRepository) FindByID(ctx context.Context, id string) (*domain.Article, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Article), args.Error(1)
}

func (m *MockArticleRepository) FindAll(ctx context.Context) ([]*domain.Article, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*domain.Article), args.Error(1)
}
