package repository_test

import (
	"context"
	"testing"

	"github.com/dynastymasra/mujib/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/dynastymasra/mujib/config"
	"github.com/dynastymasra/mujib/infrastructure/provider/postgres"
	"github.com/dynastymasra/mujib/infrastructure/repository"
	"github.com/stretchr/testify/suite"
)

type ArticleRepositorySuite struct {
	suite.Suite
	*repository.ArticleRepository
}

func Test_ArticleRepository(t *testing.T) {
	suite.Run(t, new(ArticleRepositorySuite))
}

func (a *ArticleRepositorySuite) SetupSuite() {
	config.Load()
	config.SetupTestLogger()
}

func (a *ArticleRepositorySuite) SetupTest() {
	db, _ := postgres.Connect(config.Postgres())

	articleRepo := repository.NewArticleRepository(db)

	a.ArticleRepository = articleRepo
}

func (a *ArticleRepositorySuite) TearDownSuite() {
	db, _ := postgres.Connect(config.Postgres())
	postgres.Close(db)
	postgres.Reset()
}

func article() domain.Article {
	return domain.Article{
		ID:      uuid.NewV4().String(),
		Title:   "Hello World",
		Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		Author:  "John",
	}
}

func (a *ArticleRepositorySuite) Test_Save_Success() {
	err := a.ArticleRepository.Save(context.Background(), article())

	assert.NoError(a.T(), err)
}

func (a *ArticleRepositorySuite) Test_Save_Failed() {
	article := domain.Article{
		ID:      "test",
		Title:   "Hello World",
		Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		Author:  "John",
	}

	err := a.ArticleRepository.Save(context.Background(), article)

	assert.Error(a.T(), err)
}

func (a *ArticleRepositorySuite) Test_FindByID_Success() {
	article := article()

	a.ArticleRepository.Save(context.Background(), article)

	resp, err := a.ArticleRepository.FindByID(context.Background(), article.ID)

	assert.NotNil(a.T(), resp)
	assert.NoError(a.T(), err)
}

func (a *ArticleRepositorySuite) Test_FindByID_Failed_NotFound() {
	resp, err := a.ArticleRepository.FindByID(context.Background(), uuid.NewV4().String())

	assert.Empty(a.T(), resp)
	assert.Error(a.T(), err)
}

func (a *ArticleRepositorySuite) Test_FindAll_Success() {
	a.ArticleRepository.Save(context.Background(), article())

	articles, err := a.ArticleRepository.FindAll(context.Background())

	assert.NotEmpty(a.T(), articles)
	assert.NotNil(a.T(), articles)
	assert.NotZero(a.T(), articles)
	assert.NoError(a.T(), err)
}
