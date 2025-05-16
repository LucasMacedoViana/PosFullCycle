package entity

import (
	"apis/pkg/entity"
	"errors"
	"time"
)

var (
	ErrIdIsRequired  = errors.New("id is required")
	ErrInvalidId     = errors.New("invalid id")
	ErrNameRequired  = errors.New("name is required")
	ErrPriceInvalid  = errors.New("price is invalid")
	ErrPriceRequired = errors.New("price is required")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIdIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidId
	}
	if p.Name == "" {
		return ErrNameRequired
	}
	if p.Price == 0 {
		return ErrPriceRequired
	}
	if p.Price < 0 {
		return ErrPriceInvalid
	}
	return nil
}
