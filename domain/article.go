package domain

import "context"

type Article struct {
	ID      string `json:"id" gorm:"not null;column:id;primary_key" validate:"omitempty,uuid"`
	Title   string `json:"title" gorm:"not null;column:title" validate:"required,max=10"`
	Content string `json:"content" gorm:"not null;column:content" validate:"required"`
	Author  string `json:"author" gorm:"not null;column:author" validate:"required,max=50"`
}

func (Article) TableName() string {
	return "articles"
}

type ArticleRepository interface {
	Save(context.Context, Article) error
	FindByID(context.Context, string) (*Article, error)
	FindAll(context.Context) ([]*Article, error)
}
