package domain

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type ArrayString []string

type Product struct {
	ID                    int         `json:"product_id" gorm:"column:id;primary_key" validate:"omitempty"`
	Name                  string      `json:"name" gorm:"column:name;" validate:"required,max=255"`
	ImageClosed           string      `json:"image_closed" gorm:"column:image_closed" validate:"required"`
	ImageOpen             string      `json:"image_open" gorm:"column:image_open" validate:"required"`
	Description           string      `json:"description" gorm:"column:description" validate:"required"`
	Story                 string      `json:"story" gorm:"column:story" validate:""`
	SourcingValues        ArrayString `json:"sourcing_values" gorm:"column:sourcing_values" validate:"dive,required"`
	Ingredients           ArrayString `json:"ingredients" gorm:"column:ingredients" validate:"dive,required"`
	AllergyInfo           string      `json:"allergy_info" gorm:"column:allergy_info" validate:"omitempty"`
	DietaryCertifications string      `json:"dietary_certifications" gorm:"column:dietary_certifications" validate:"omitempty"`
}

func (Product) TableName() string {
	return "products"
}

func (a ArrayString) Value() (driver.Value, error) {
	value, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *ArrayString) Scan(value interface{}) error {
	if err := json.Unmarshal([]byte(fmt.Sprint(value)), &a); err != nil {
		return err
	}

	return nil
}

type ProductRepository interface {
	Save(context.Context, Product) error
	FindByID(context.Context, int) (*Product, error)
	Fetch(context.Context) ([]Product, error)
	Update(context.Context, Product) error
	Delete(context.Context, Product) error
}

type ProductService interface {
	Create(context.Context, Product) (*Product, error)
	Delete(context.Context, string) error
}
