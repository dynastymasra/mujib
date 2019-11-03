package handler_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/dynastymasra/mujib/config"
	"github.com/dynastymasra/mujib/delivery/http/handler"
	"github.com/dynastymasra/mujib/test"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ProductDeleteSuite struct {
	suite.Suite
	productService *test.MockProductService
}

func Test_ProductDeleteSuite(t *testing.T) {
	suite.Run(t, new(ProductDeleteSuite))
}

func (p *ProductDeleteSuite) SetupSuite() {
	config.SetupTestLogger()
}

func (p *ProductDeleteSuite) SetupTest() {
	p.productService = &test.MockProductService{}
}

func (p *ProductDeleteSuite) Test_ProductDelete_Success() {
	id := uuid.NewV4().String()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/products/%s", id), nil)

	r = mux.SetURLVars(r, map[string]string{
		"product_id": id,
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("Delete", ctx, id).Return(nil)

	handler.ProductDelete(p.productService)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusOK, w.Code)
}

func (p *ProductDeleteSuite) Test_ProductDelete_Failed_NotFound() {
	id := uuid.NewV4().String()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/products/%s", id), nil)

	r = mux.SetURLVars(r, map[string]string{
		"product_id": id,
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("Delete", ctx, id).Return(gorm.ErrRecordNotFound)

	handler.ProductDelete(p.productService)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusNotFound, w.Code)
}

func (p *ProductDeleteSuite) Test_ProductDelete_Failed() {
	id := uuid.NewV4().String()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/products/%s", id), nil)

	r = mux.SetURLVars(r, map[string]string{
		"product_id": id,
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("Delete", ctx, id).Return(assert.AnError)

	handler.ProductDelete(p.productService)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusInternalServerError, w.Code)
}
