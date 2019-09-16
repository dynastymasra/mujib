package service_test

import (
	"context"
	"errors"
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/dynastymasra/mujib/config"
	"github.com/dynastymasra/mujib/domain"
	"github.com/dynastymasra/mujib/service"
	"github.com/dynastymasra/mujib/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ArticleServiceSuite struct {
	suite.Suite
	articleRepo    *test.MockArticleRepository
	articleService *service.ArticleService
}

func Test_ArticleServiceSuite(t *testing.T) {
	suite.Run(t, new(ArticleServiceSuite))
}

func (a *ArticleServiceSuite) SetupSuite() {
	config.SetupTestLogger()
}

func (a *ArticleServiceSuite) SetupTest() {
	a.articleRepo = &test.MockArticleRepository{}
	articleService := service.NewArticleService(a.articleRepo)
	a.articleService = &articleService
}

func (a *ArticleServiceSuite) Test_CreateArticle_Success() {
	a.articleRepo.On("Save", context.Background()).Return(nil)

	article, err := a.articleService.CreateArticle(context.Background(), domain.Article{})

	assert.NotNil(a.T(), article)
	assert.NoError(a.T(), err)
}

func (a *ArticleServiceSuite) Test_CreateArticle_Failed() {
	a.articleRepo.On("Save", context.Background()).Return(errors.New("failed"))

	article, err := a.articleService.CreateArticle(context.Background(), domain.Article{})

	assert.Nil(a.T(), article)
	assert.Error(a.T(), err)
}

func (a *ArticleServiceSuite) Test_FindArticleByID_Success() {
	id := uuid.NewV4().String()

	a.articleRepo.On("FindByID", context.Background(), id).Return(&domain.Article{}, nil)

	article, err := a.articleService.FindArticleByID(context.Background(), id)

	assert.NotNil(a.T(), article)
	assert.NoError(a.T(), err)
}

func (a *ArticleServiceSuite) Test_FindArticleByID_Failed() {
	id := uuid.NewV4().String()

	a.articleRepo.On("FindByID", context.Background(), id).Return((*domain.Article)(nil), errors.New("failed"))

	article, err := a.articleService.FindArticleByID(context.Background(), id)

	assert.Nil(a.T(), article)
	assert.Error(a.T(), err)
}

func (a *ArticleServiceSuite) Test_FindAllArticle_Success() {
	a.articleRepo.On("FindAll", context.Background()).Return([]*domain.Article{{}, {}}, nil)

	articles, err := a.articleService.FindAllArticle(context.Background())

	assert.NotNil(a.T(), articles)
	assert.NotEmpty(a.T(), articles)
	assert.NoError(a.T(), err)
}

func (a *ArticleServiceSuite) Test_FindAllArticle_Empty() {
	a.articleRepo.On("FindAll", context.Background()).Return([]*domain.Article{}, nil)

	articles, err := a.articleService.FindAllArticle(context.Background())

	assert.NotNil(a.T(), articles)
	assert.Empty(a.T(), articles)
	assert.NoError(a.T(), err)
}

func (a *ArticleServiceSuite) Test_FindAllArticle_Failed() {
	a.articleRepo.On("FindAll", context.Background()).Return(([]*domain.Article)(nil), errors.New("failed"))

	articles, err := a.articleService.FindAllArticle(context.Background())

	assert.Nil(a.T(), articles)
	assert.Error(a.T(), err)
}
