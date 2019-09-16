package article_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/dynastymasra/mujib/domain"
	"github.com/dynastymasra/mujib/infrastructure/web/controller/article"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/dynastymasra/mujib/config"
	"github.com/dynastymasra/mujib/test"
	"github.com/stretchr/testify/suite"
)

type FindSuite struct {
	suite.Suite
	articleService *test.MockArticleService
}

func Test_FindSuite(t *testing.T) {
	suite.Run(t, new(FindSuite))
}

func (f *FindSuite) SetupSuite() {
	config.SetupTestLogger()
}

func (f *FindSuite) SetupTest() {
	f.articleService = &test.MockArticleService{}
}

func (f *FindSuite) Test_FindByID_Success() {
	id := uuid.NewV4().String()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/articles/%s", id), nil)

	r = mux.SetURLVars(r, map[string]string{
		"article_id": id,
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	f.articleService.On("FindArticleByID", ctx, id).Return(&domain.Article{}, nil)

	article.FindByID(f.articleService)(w, r.WithContext(ctx))

	assert.Equal(f.T(), http.StatusOK, w.Code)
}

func (f *FindSuite) Test_FindByID_NotFound() {
	id := uuid.NewV4().String()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/articles/%s", id), nil)

	r = mux.SetURLVars(r, map[string]string{
		"article_id": id,
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	f.articleService.On("FindArticleByID", ctx, id).Return((*domain.Article)(nil), gorm.ErrRecordNotFound)

	article.FindByID(f.articleService)(w, r.WithContext(ctx))

	assert.Equal(f.T(), http.StatusNotFound, w.Code)
}

func (f *FindSuite) Test_FindByID_Failed() {
	id := uuid.NewV4().String()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/articles/%s", id), nil)

	r = mux.SetURLVars(r, map[string]string{
		"article_id": id,
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	f.articleService.On("FindArticleByID", ctx, id).Return((*domain.Article)(nil), errors.New("failed"))

	article.FindByID(f.articleService)(w, r.WithContext(ctx))

	assert.Equal(f.T(), http.StatusInternalServerError, w.Code)
}

func (f *FindSuite) Test_FindAll_Success() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/articles", nil)

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	f.articleService.On("FindAllArticle", ctx).Return([]*domain.Article{{}, {}}, nil)

	article.FindAll(f.articleService)(w, r.WithContext(ctx))

	assert.Equal(f.T(), http.StatusOK, w.Code)
}

func (f *FindSuite) Test_FindAll_Failed() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/articles", nil)

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	f.articleService.On("FindAllArticle", ctx).Return(([]*domain.Article)(nil), errors.New("failed"))

	article.FindAll(f.articleService)(w, r.WithContext(ctx))

	assert.Equal(f.T(), http.StatusInternalServerError, w.Code)
}
