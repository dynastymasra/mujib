package product

import (
	"context"

	"github.com/dynastymasra/mujib/config"
	"github.com/dynastymasra/mujib/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type Service struct {
	ProductRepository domain.ProductRepository
}

func NewService(productRepo domain.ProductRepository) Service {
	return Service{ProductRepository: productRepo}
}

func (s Service) Create(ctx context.Context, product domain.Product) (*domain.Product, error) {
	if len(product.ID) < 1 {
		product.ID = uuid.NewV4().String()
	}

	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
		"product":        product,
	})

	if err := s.ProductRepository.Save(ctx, product); err != nil {
		log.WithError(err).Errorln("Failed create new product")
		return nil, err
	}

	return &product, nil
}

func (s Service) Update(ctx context.Context, id string, product domain.Product) (*domain.Product, error) {
	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
		"after":          product,
		"id":             id,
	})

	prod, err := s.ProductRepository.FindByID(ctx, id)
	if err != nil {
		log.WithError(err).Errorln("Failed get product by id")
		return nil, err
	}

	product.ID = id
	if err := s.ProductRepository.Update(ctx, product); err != nil {
		log.WithField("before", prod).WithError(err).Errorln("Failed update product")
		return nil, err
	}

	return &product, nil
}

func (s Service) Delete(ctx context.Context, id string) error {
	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
		"id":             id,
	})

	product, err := s.ProductRepository.FindByID(ctx, id)
	if err != nil {
		log.WithError(err).Errorln("Failed get product")
		return err
	}

	if err := s.ProductRepository.Delete(ctx, *product); err != nil {
		log.WithError(err).Errorln("Failed delete product")
		return err
	}

	return nil
}
