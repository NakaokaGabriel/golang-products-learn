package database

import (
	"testing"

	internal "github.com/NakaokaGabriel/go/products-projects/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&internal.User{})

	user, _ := internal.NewUser("John", "john@teste.com", "12345")
	userDb := NewUser(db)

	err = userDb.Create(user)

	assert.Nil(t, err)

	var findUserById *internal.User

	err = db.Where("id = ?", user.Id).First(&findUserById).Error

	assert.Nil(t, err)
	assert.Equal(t, "John", findUserById.Name)
	assert.Equal(t, "john@teste.com", findUserById.Email)
	assert.Equal(t, user.Id, findUserById.Id)
	assert.NotNil(t, findUserById)
}

func TestFindUserByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&internal.User{})

	user, _ := internal.NewUser("John", "john@teste.com", "12345")
	userDb := NewUser(db)

	err = userDb.Create(user)
	assert.Nil(t, err)

	userFound, err := userDb.FindByEmail(user.Email)

	assert.Nil(t, err)
	assert.NotNil(t, userFound)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Name, userFound.Name)
}
