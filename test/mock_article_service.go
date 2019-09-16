package test

import (
	"context"

	"github.com/dynastymasra/mujib/domain"
	"github.com/stretchr/testify/mock"
)

type MockArticleService struct {
	mock.Mock
}

func (m *MockArticleService) CreateArticle(ctx context.Context, article domain.Article) (*domain.Article, error) {
	args := m.Called(ctx)
	return args.Get(0).(*domain.Article), args.Error(1)
}

func (m *MockArticleService) FindArticleByID(ctx context.Context, id string) (*domain.Article, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Article), args.Error(1)
}

func (m *MockArticleService) FindAllArticle(ctx context.Context) ([]*domain.Article, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*domain.Article), args.Error(1)
}
