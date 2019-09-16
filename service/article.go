package service

import (
	"context"

	"github.com/dynastymasra/mujib/config"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	"github.com/dynastymasra/mujib/domain"
	"github.com/dynastymasra/mujib/infrastructure/repository"
)

type ArticleServicer interface {
	CreateArticle(context.Context, domain.Article) (*domain.Article, error)
}

type ArticleService struct {
	ArticleRepository *repository.ArticleRepository
}

func NewArticleService(articleRepo *repository.ArticleRepository) ArticleService {
	return ArticleService{
		ArticleRepository: articleRepo,
	}
}

func (a ArticleService) CreateArticle(ctx context.Context, article domain.Article) (*domain.Article, error) {
	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
		"article":        article,
	})

	article.ID = uuid.NewV4().String()

	if err := a.ArticleRepository.Save(ctx, article); err != nil {
		log.WithError(err).Errorln("Failed create new article")
		return nil, err
	}

	result, err := a.ArticleRepository.FindByID(ctx, article.ID)
	if err != nil {
		log.WithError(err).Errorln("Failed get article from db")
		return nil, err
	}

	return result, nil
}
