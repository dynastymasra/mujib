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
	FindArticleByID(context.Context, string) (*domain.Article, error)
	FindAllArticle(context.Context) ([]*domain.Article, error)
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

	return &article, nil
}

func (a ArticleService) FindArticleByID(ctx context.Context, id string) (*domain.Article, error) {
	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
		"article_id":     id,
	})

	article, err := a.ArticleRepository.FindByID(ctx, id)
	if err != nil {
		log.WithError(err).Errorln("Failed find by id article")
		return nil, err
	}

	return article, nil
}

func (a ArticleService) FindAllArticle(ctx context.Context) ([]*domain.Article, error) {
	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
	})

	articles, err := a.ArticleRepository.FindAll(ctx)
	if err != nil {
		log.WithError(err).Errorln("Failed find all article")
		return nil, err
	}

	return articles, nil
}
