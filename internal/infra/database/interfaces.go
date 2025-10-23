package database

import internal "github.com/NakaokaGabriel/go/products-projects/internal/entity"

type UserInterface interface {
	Create(user *internal.User) error
	FindByEmail(email string) (*internal.User, error)
}
