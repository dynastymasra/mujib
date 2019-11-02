package domain

import "context"

type Product struct {
	ID                    string   `json:"product_id" gorm:"not null;column:id;primary_key" validate:"omitempty"`
	Name                  string   `json:"name" gorm:"not null;column:name;" validate:"required,max=255"`
	ImageClosed           string   `json:"image_closed" gorm:"not null;column:image_closed" validate:"required"`
	ImageOpen             string   `json:"image_open" gorm:"not null;column:image_open" validate:"required"`
	Description           string   `json:"description" gorm:"not null;column:description" validate:"required"`
	Story                 string   `json:"story" gorm:"not null;column:story" validate:"required"`
	SourcingValues        []string `json:"sourcing_values" gorm:"not null;column:sourcing_values" validate:"dive,required"`
	Ingredients           []string `json:"ingredients" gorm:"not null;column:ingredients" validate:"dive,required"`
	AllergyInfo           string   `json:"allergy_info" gorm:"not null;column:allergy_info" validate:"required,max=255"`
	DietaryCertifications string   `json:"dietary_certifications" gorm:"not null;column:dietary_certifications" validate:"required,max=50"`
}

func (Product) TableName() string {
	return "products"
}

type ProductRepository interface {
	Save(context.Context, Product) error
	FindByID(context.Context, string) (*Article, error)
	FindAll(context.Context) ([]Article, error)
	Update(context.Context, Product) error
	Delete(context.Context, string) error
}
