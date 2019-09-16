package article_test

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dynastymasra/mujib/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/dynastymasra/mujib/infrastructure/web/controller/article"

	"github.com/dynastymasra/mujib/config"
	"github.com/dynastymasra/mujib/test"
	"github.com/stretchr/testify/suite"
)

type SaveSuite struct {
	suite.Suite
	articleService *test.MockArticleService
}

func Test_SaveSuite(t *testing.T) {
	suite.Run(t, new(SaveSuite))
}

func (s *SaveSuite) SetupSuite() {
	config.SetupTestLogger()
}

func (s *SaveSuite) SetupTest() {
	s.articleService = &test.MockArticleService{}
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("failed")
}

func (s *SaveSuite) Test_Save_Success() {
	reqInBytes := []byte(`{
		"title": "Hello World",
		"content": "Lorem",
		"author": "John"
	}`)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewReader(reqInBytes))

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	s.articleService.On("CreateArticle", ctx).Return(&domain.Article{}, nil)

	article.Save(s.articleService)(w, r.WithContext(ctx))
	assert.Equal(s.T(), http.StatusCreated, w.Code)
}

func (s *SaveSuite) Test_Save_FailedReadBody() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/articles", errReader(0))

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	article.Save(s.articleService)(w, r.WithContext(ctx))
	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
}

func (s *SaveSuite) Test_Save_FailedUnmarshal() {
	reqInBytes := []byte(`<- test chan`)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewReader(reqInBytes))

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	article.Save(s.articleService)(w, r.WithContext(ctx))
	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
}

func (s *SaveSuite) Test_Save_FailedValidation() {
	reqInBytes := []byte(`{
		"title": "Hello World",
		"content": "Lorem",
		"author": "123456789012345678901234567890123456789012345678901234567890"
	}`)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewReader(reqInBytes))

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	article.Save(s.articleService)(w, r.WithContext(ctx))
	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
}

func (s *SaveSuite) Test_Save_Failed() {
	reqInBytes := []byte(`{
		"title": "Hello World",
		"content": "Lorem",
		"author": "John"
	}`)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewReader(reqInBytes))

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	s.articleService.On("CreateArticle", ctx).Return((*domain.Article)(nil), errors.New("failed"))

	article.Save(s.articleService)(w, r.WithContext(ctx))
	assert.Equal(s.T(), http.StatusInternalServerError, w.Code)
}
