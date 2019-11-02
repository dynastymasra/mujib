package product_test

import (
	"context"
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/assert"

	"github.com/dynastymasra/mujib/domain"

	"github.com/dynastymasra/mujib/config"

	"github.com/dynastymasra/mujib/product"
	"github.com/dynastymasra/mujib/test"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
	productRepo    *test.MockProductRepository
	productService *product.Service
}

func Test_ServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) SetupSuite() {
	config.SetupTestLogger()
}

func (s *ServiceSuite) SetupTest() {
	s.productRepo = &test.MockProductRepository{}
	productService := product.NewService(s.productRepo)
	s.productService = &productService
}

func (s *ServiceSuite) Test_Create_Success() {
	s.productRepo.On("Save", context.Background()).Return(nil)

	product, err := s.productService.Create(context.Background(), domain.Product{})

	assert.NotNil(s.T(), product)
	assert.NoError(s.T(), err)
}

func (s *ServiceSuite) Test_Create_Failed() {
	s.productRepo.On("Save", context.Background()).Return(assert.AnError)

	product, err := s.productService.Create(context.Background(), domain.Product{})

	assert.Nil(s.T(), product)
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) Test_Update_Success() {
	id := uuid.NewV4().String()
	s.productRepo.On("FindByID", context.Background(), id).Return(&domain.Product{ID: id}, nil)
	s.productRepo.On("Update", context.Background(), domain.Product{ID: id}).Return(nil)

	product, err := s.productService.Update(context.Background(), id, domain.Product{})

	assert.NotNil(s.T(), product)
	assert.Equal(s.T(), id, product.ID)
	assert.NoError(s.T(), err)
}

func (s *ServiceSuite) Test_Update_Failed_FindByID() {
	id := uuid.NewV4().String()
	s.productRepo.On("FindByID", context.Background(), id).Return((*domain.Product)(nil), assert.AnError)

	product, err := s.productService.Update(context.Background(), id, domain.Product{})

	assert.Nil(s.T(), product)
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) Test_Update_Failed_Update() {
	id := uuid.NewV4().String()

	s.productRepo.On("FindByID", context.Background(), id).Return(&domain.Product{ID: id}, nil)
	s.productRepo.On("Update", context.Background(), domain.Product{ID: id}).Return(assert.AnError)

	product, err := s.productService.Update(context.Background(), id, domain.Product{})

	assert.Nil(s.T(), product)
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) Test_Delete_Success() {
	id := uuid.NewV4().String()

	s.productRepo.On("FindByID", context.Background(), id).Return(&domain.Product{ID: id}, nil)
	s.productRepo.On("Delete", context.Background(), domain.Product{ID: id}).Return(nil)

	err := s.productService.Delete(context.Background(), id)

	assert.NoError(s.T(), err)
}

func (s *ServiceSuite) Test_Delete_Failed_FindByID() {
	id := uuid.NewV4().String()

	s.productRepo.On("FindByID", context.Background(), id).Return((*domain.Product)(nil), assert.AnError)

	err := s.productService.Delete(context.Background(), id)

	assert.Error(s.T(), err)
}

func (s *ServiceSuite) Test_Delete_Failed_Delete() {
	id := uuid.NewV4().String()

	s.productRepo.On("FindByID", context.Background(), id).Return(&domain.Product{ID: id}, nil)
	s.productRepo.On("Delete", context.Background(), domain.Product{ID: id}).Return(assert.AnError)

	err := s.productService.Delete(context.Background(), id)

	assert.Error(s.T(), err)
}
