package database

import internal "github.com/NakaokaGabriel/go/products-projects/internal/entity"

type UserInterface interface {
	Create(user *internal.User) error
	FindByEmail(email string) (*internal.User, error)
}

type ProductInterface interface {
	Create(product *internal.Product) error
	Update(product *internal.Product) error
	Delete(id string) error
	FindAll(page int, limit int, sort string) ([]internal.Product, error)
	FindById(id string) (*internal.Product, error)
}
