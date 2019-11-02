package handler_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/dynastymasra/mujib/delivery/http/handler"
	"github.com/dynastymasra/mujib/domain"
	"github.com/stretchr/testify/assert"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"

	"github.com/dynastymasra/mujib/config"
	"github.com/dynastymasra/mujib/test"
	"github.com/stretchr/testify/suite"
)

type ProductFetchSuite struct {
	suite.Suite
	productService *test.MockProductService
}

func Test_ProductFetchSuite(t *testing.T) {
	suite.Run(t, new(ProductFetchSuite))
}

func (p *ProductFetchSuite) SetupSuite() {
	config.SetupTestLogger()
}

func (p *ProductFetchSuite) SetupTest() {
	p.productService = &test.MockProductService{}
}

func (p *ProductFetchSuite) Test_ProductFindByID_Success() {
	id := uuid.NewV4().String()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/products/%s", id), nil)

	r = mux.SetURLVars(r, map[string]string{
		"product_id": id,
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("FindByID", ctx, id).Return(&domain.Product{}, nil)

	handler.ProductFindByID(p.productService)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusOK, w.Code)
}

func (p *ProductFetchSuite) Test_ProductFindByID_Failed_NotFound() {
	id := uuid.NewV4().String()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/products/%s", id), nil)

	r = mux.SetURLVars(r, map[string]string{
		"product_id": id,
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("FindByID", ctx, id).Return((*domain.Product)(nil), gorm.ErrRecordNotFound)

	handler.ProductFindByID(p.productService)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusNotFound, w.Code)
}

func (p *ProductFetchSuite) Test_ProductFindByID_Failed() {
	id := uuid.NewV4().String()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/products/%s", id), nil)

	r = mux.SetURLVars(r, map[string]string{
		"product_id": id,
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("FindByID", ctx, id).Return((*domain.Product)(nil), assert.AnError)

	handler.ProductFindByID(p.productService)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusInternalServerError, w.Code)
}

func (p *ProductFetchSuite) Test_ProductFindAll_Success() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/products?from=20&size=40", nil)

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("Fetch", ctx, 20, 40).Return([]domain.Product{{ID: "id"}}, nil)

	handler.ProductFindAll(p.productService)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusOK, w.Code)
}

func (p *ProductFetchSuite) Test_ProductFindAll_Failed() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/products?from=20&size=40", nil)

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("Fetch", ctx, 20, 40).Return(([]domain.Product)(nil), assert.AnError)

	handler.ProductFindAll(p.productService)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusInternalServerError, w.Code)
}
