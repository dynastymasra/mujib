package handler_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/dynastymasra/mujib/delivery/http/handler"
	"github.com/dynastymasra/mujib/domain"
	uuid "github.com/satori/go.uuid"

	"github.com/dynastymasra/mujib/config"
	"github.com/dynastymasra/mujib/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ProductUpdateSuite struct {
	suite.Suite
	productService *test.MockProductService
}

func Test_ProductUpdateSuite(t *testing.T) {
	suite.Run(t, new(ProductUpdateSuite))
}

func (p *ProductUpdateSuite) SetupSuite() {
	config.SetupTestLogger()
}

func (p *ProductUpdateSuite) SetupTest() {
	p.productService = &test.MockProductService{}
}

func (p *ProductCreateSuite) Test_ProductUpdate_Success() {
	id := uuid.NewV4().String()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/v1/products/%s", id), bytes.NewReader(productPayload()))
	r = mux.SetURLVars(r, map[string]string{
		"product_id": id,
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("Update", ctx, id).Return(&domain.Product{}, nil)

	handler.ProductUpdate(p.productService)(w, r.WithContext(ctx))
	assert.Equal(p.T(), http.StatusOK, w.Code)
}

func (p *ProductCreateSuite) Test_ProductUpdate_Failed_ReadBody() {
	id := uuid.NewV4().String()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/v1/products/%s", id), errReader(0))
	r = mux.SetURLVars(r, map[string]string{
		"product_id": id,
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	handler.ProductUpdate(p.productService)(w, r.WithContext(ctx))
	assert.Equal(p.T(), http.StatusBadRequest, w.Code)
}

func (p *ProductCreateSuite) Test_ProductUpdate_Failed_Unmarshal() {
	reqInBytes := []byte(`<- test chan`)
	id := uuid.NewV4().String()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/v1/products/%s", id), bytes.NewReader(reqInBytes))
	r = mux.SetURLVars(r, map[string]string{
		"product_id": id,
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	handler.ProductUpdate(p.productService)(w, r.WithContext(ctx))
	assert.Equal(p.T(), http.StatusBadRequest, w.Code)
}

func (p *ProductCreateSuite) Test_ProductUpdate_Failed_Validation() {
	reqInBytes := []byte(`{
		"name": "test"
	}`)
	id := uuid.NewV4().String()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/v1/products/%s", id), bytes.NewReader(reqInBytes))
	r = mux.SetURLVars(r, map[string]string{
		"product_id": id,
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	handler.ProductUpdate(p.productService)(w, r.WithContext(ctx))
	assert.Equal(p.T(), http.StatusBadRequest, w.Code)
}

func (p *ProductCreateSuite) Test_ProductUpdate_Failed() {
	id := uuid.NewV4().String()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/v1/products/%s", id), bytes.NewReader(productPayload()))
	r = mux.SetURLVars(r, map[string]string{
		"product_id": id,
	})

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("Update", ctx, id).Return((*domain.Product)(nil), assert.AnError)

	handler.ProductUpdate(p.productService)(w, r.WithContext(ctx))
	assert.Equal(p.T(), http.StatusInternalServerError, w.Code)
}
