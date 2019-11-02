package handler_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dynastymasra/mujib/delivery/http/handler"
	"github.com/dynastymasra/mujib/domain"
	uuid "github.com/satori/go.uuid"

	"github.com/dynastymasra/mujib/config"
	"github.com/dynastymasra/mujib/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ProductCreateSuite struct {
	suite.Suite
	productService *test.MockProductService
}

func Test_SaveSuite(t *testing.T) {
	suite.Run(t, new(ProductCreateSuite))
}

func (p *ProductCreateSuite) SetupSuite() {
	config.SetupTestLogger()
}

func (p *ProductCreateSuite) SetupTest() {
	p.productService = &test.MockProductService{}
}

type errReader int

func (errReader) Read([]byte) (n int, err error) {
	return 0, assert.AnError
}

func productPayload() []byte {
	return []byte(`{
		"name": "Vanilla Toffee Bar Crunch",
		"image_closed": "/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing.png",
		"image_open": "/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing-open.png",
		"description": "Vanilla Ice Cream with Fudge-Covered Toffee Pieces",
		"story": "Vanilla What Bar Crunch? We gave this flavor a new name to go with the new toffee bars we’re using as part of our commitment to source Fairtrade Certified and non-GMO ingredients. We love it and know you will too!",
		"sourcing_values": [
			"Non-GMO",
			"Cage-Free Eggs",
			"Fairtrade",
			"Responsibly Sourced Packaging",
			"Caring Dairy"
		],
		"ingredients": [
			"cream",
			"skim milk",
			"liquid sugar",
			"water",
			"sugar",
			"coconut oil",
			"egg yolks",
			"butter",
			"vanilla extract",
			"almonds",
			"cocoa (processed with alkali)",
			"milk",
			"soy lecithin",
			"cocoa",
			"natural flavor",
			"salt",
			"vegetable oil",
			"guar gum",
			"carrageenan"
		],
		"allergy_info": "may contain wheat, peanuts and other tree nuts",
		"dietary_certifications": "Kosher"
	}`)
}

func (p *ProductCreateSuite) Test_ProductCreate_Success() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/products", bytes.NewReader(productPayload()))

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("Create", ctx).Return(&domain.Product{}, nil)

	handler.ProductCreate(p.productService)(w, r.WithContext(ctx))
	assert.Equal(p.T(), http.StatusCreated, w.Code)
}

func (p *ProductCreateSuite) Test_ProductCreate_Failed_ReadBody() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/products", errReader(0))

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	handler.ProductCreate(p.productService)(w, r.WithContext(ctx))
	assert.Equal(p.T(), http.StatusBadRequest, w.Code)
}

func (p *ProductCreateSuite) Test_ProductCreate_Failed_Unmarshal() {
	reqInBytes := []byte(`<- test chan`)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/products", bytes.NewReader(reqInBytes))

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	handler.ProductCreate(p.productService)(w, r.WithContext(ctx))
	assert.Equal(p.T(), http.StatusBadRequest, w.Code)
}

func (p *ProductCreateSuite) Test_ProductCreate_Failed_Validation() {
	reqInBytes := []byte(`{
		"name": "test"
	}`)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/products", bytes.NewReader(reqInBytes))

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	handler.ProductCreate(p.productService)(w, r.WithContext(ctx))
	assert.Equal(p.T(), http.StatusBadRequest, w.Code)
}

func (p *ProductCreateSuite) Test_ProductCreate_Failed() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/products", bytes.NewReader(productPayload()))

	ctx := context.WithValue(r.Context(), config.HeaderRequestID, uuid.NewV4().String())

	p.productService.On("Create", ctx).Return((*domain.Product)(nil), assert.AnError)

	handler.ProductCreate(p.productService)(w, r.WithContext(ctx))
	assert.Equal(p.T(), http.StatusInternalServerError, w.Code)
}
