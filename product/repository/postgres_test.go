package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	uuid "github.com/satori/go.uuid"

	"github.com/dynastymasra/mujib/domain"
	"github.com/dynastymasra/mujib/product/repository"

	"github.com/dynastymasra/mujib/config"
	"github.com/dynastymasra/mujib/infrastructure/database/postgres"
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
	db, _ := postgres.Connect(config.Postgres())
	postgres.Close(db)
	postgres.Reset()
}

func product() domain.Product {
	return domain.Product{
		ID:          uuid.NewV4().String(),
		Name:        "Vanilla Toffee Bar Crunch",
		ImageClosed: "/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing.png",
		ImageOpen:   "/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing-open.png",
		Description: "Vanilla Ice Cream with Fudge-Covered Toffee Pieces",
		Story:       "Vanilla What Bar Crunch? We gave this flavor a new name to go with the new toffee bars we’re using as part of our commitment to source Fairtrade Certified and non-GMO ingredients. We love it and know you will too!",
		SourcingValues: []string{
			"Non-GMO",
			"Cage-Free Eggs",
			"Fairtrade",
			"Responsibly Sourced Packaging",
			"Caring Dairy",
		},
		Ingredients: []string{"cream",
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
