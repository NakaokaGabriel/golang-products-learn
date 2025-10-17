package user

import (
	id "github.com/NakaokaGabriel/go/products-projects/pkg/entity"
	bycrypt "golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       id.ID  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func NewUser(name string, email string, password string) (*User, error) {
	hashedPassword, err := bycrypt.GenerateFromPassword([]byte(password), bycrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return &User{
		Id:       id.NewId(),
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}, nil
}

func (u *User) ComparePassword(password string) bool {
	err := bycrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}
