package database

import (
	internal "github.com/NakaokaGabriel/go/products-projects/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *internal.User) error {
	err := u.DB.Create(user).Error

	return err
}

func (u *User) FindByEmail(email string) (*internal.User, error) {
	var user internal.User

	err := u.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
