package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Iphone", 10)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.Id)
	assert.Equal(t, "Iphone", product.Name)
	assert.Equal(t, 10, product.Price)
}

func TestProductNameRequired(t *testing.T) {
	product, err := NewProduct("", 10)

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "Name is Required")
}

func TestProductPriceRequired(t *testing.T) {
	product, err := NewProduct("Iphone", 0)

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.EqualError(t, ErrorPriceIsRequired, err.Error())
}

func TestProductInvalidPrice(t *testing.T) {
	product, err := NewProduct("Iphone", -10)

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.EqualError(t, errorPriceIsInvalid, err.Error())
}
