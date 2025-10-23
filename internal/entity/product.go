package internal

import (
	"errors"
	"time"

	"github.com/NakaokaGabriel/go/products-projects/pkg/entity"
)

type Product struct {
	Id        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	ErrorIdRequired      = errors.New("Id is required")
	ErrorIdIsInvalid     = errors.New("Id is invalid")
	ErrorNameRequired    = errors.New("Name is Required")
	ErrorPriceIsRequired = errors.New("Price is Required")
	errorPriceIsInvalid  = errors.New("Price is Invalid")
)

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		Id:        entity.NewId(),
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
	if p.Id.String() == "" {
		return ErrorIdRequired
	}

	if _, err := entity.ParseId(p.Id.String()); err != nil {
		return ErrorIdIsInvalid
	}

	if p.Name == "" {
		return ErrorNameRequired
	}

	if p.Price == 0 {
		return ErrorPriceIsRequired
	}

	if p.Price < 0 {
		return errorPriceIsInvalid
	}

	return nil
}
