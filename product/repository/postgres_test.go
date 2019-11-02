package repository_test

import (
	"context"
	"log"
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/assert"

	"github.com/dynastymasra/mujib/domain"
	"github.com/dynastymasra/mujib/product/repository"

	"github.com/dynastymasra/mujib/config"
	database "github.com/dynastymasra/mujib/infrastructure/database/postgres"
	"github.com/stretchr/testify/suite"
)

type ProductRepositorySuite struct {
	suite.Suite
	*repository.ProductRepository
}

func Test_ProductRepositorySuite(t *testing.T) {
	suite.Run(t, new(ProductRepositorySuite))
}

func (p *ProductRepositorySuite) SetupSuite() {
	config.Load()
	config.SetupTestLogger()
}

func (p *ProductRepositorySuite) TearDownSuite() {
	db, _ := database.Connect(config.Postgres())
	database.Close(db)
	database.Reset()
}

func (p *ProductRepositorySuite) SetupTest() {
	db, err := database.Connect(config.Postgres())
	if err != nil {
		log.Fatal(err)
	}

	productRepo := repository.NewProductRepository(db)

	p.ProductRepository = productRepo
}

func product() domain.Product {
	return domain.Product{
		ID:          uuid.NewV4().String(),
		Name:        "Vanilla Toffee Bar Crunch",
		ImageClosed: "/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing.png",
		ImageOpen:   "/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing-open.png",
		Description: "Vanilla Ice Cream with Fudge-Covered Toffee Pieces",
		Story:       "Vanilla What Bar Crunch? We gave this flavor a new name to go with the new toffee bars we’re using as part of our commitment to source Fairtrade Certified and non-GMO ingredients. We love it and know you will too!",
		SourcingValues: domain.ArrayString{
			"Non-GMO",
			"Cage-Free Eggs",
			"Fairtrade",
			"Responsibly Sourced Packaging",
			"Caring Dairy",
		},
		Ingredients: domain.ArrayString{
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
			"carrageenan"},
		AllergyInfo:           "may contain wheat, peanuts and other tree nuts",
		DietaryCertifications: "Kosher",
	}
}

func (p *ProductRepositorySuite) Test_Save_Success() {
	err := p.ProductRepository.Save(context.Background(), product())

	assert.NoError(p.T(), err)
}

func (p *ProductRepositorySuite) Test_Save_Failed() {
	testProd := product()
	testProd.DietaryCertifications = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

	err := p.ProductRepository.Save(context.Background(), testProd)

	assert.Error(p.T(), err)
}

func (p *ProductRepositorySuite) Test_FindByID_Success() {
	product := product()

	p.ProductRepository.Save(context.Background(), product)

	resp, err := p.ProductRepository.FindByID(context.Background(), product.ID)

	assert.NotNil(p.T(), resp)
	assert.NoError(p.T(), err)
}

func (p *ProductRepositorySuite) Test_FindByID_Failed() {
	resp, err := p.ProductRepository.FindByID(context.Background(), uuid.NewV4().String())

	assert.Nil(p.T(), resp)
	assert.Error(p.T(), err)
}

func (p *ProductRepositorySuite) Test_Fetch_Success() {
	product := product()

	p.ProductRepository.Save(context.Background(), product)

	resp, err := p.ProductRepository.Fetch(context.Background(), 0, 20)

	assert.NotEmpty(p.T(), resp)
	assert.NoError(p.T(), err)
}

func (p *ProductRepositorySuite) Test_Update_Success() {
	product := product()

	p.ProductRepository.Save(context.Background(), product)

	product.Name = "Update"
	err := p.ProductRepository.Update(context.Background(), product)

	assert.NoError(p.T(), err)
}

func (p *ProductRepositorySuite) Test_Update_Failed() {
	product := product()

	p.ProductRepository.Save(context.Background(), product)

	product.DietaryCertifications = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	err := p.ProductRepository.Update(context.Background(), product)

	assert.Error(p.T(), err)
}

func (p *ProductRepositorySuite) Test_Delete_Success() {
	product := product()

	p.ProductRepository.Save(context.Background(), product)

	err := p.ProductRepository.Delete(context.Background(), product)

	assert.NoError(p.T(), err)
}
