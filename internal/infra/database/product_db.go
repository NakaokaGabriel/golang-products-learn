package database

import (
	internal "github.com/NakaokaGabriel/go/products-projects/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product internal.Product) error {
	err := p.DB.Create(&product).Error

	return err
}

func (p *Product) FindAll(page int, limit int, sort string) ([]internal.Product, error) {
	var products []internal.Product
	var err error

	if sort == "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}

	return products, err
}

func (p *Product) FindById(id string) (*internal.Product, error) {
	var product internal.Product

	err := p.DB.Where("id = ?", id).First(&product).Error

	return &product, err
}

func (p *Product) Update(product *internal.Product) error {
	_, err := p.FindById(product.Id.String())

	if err != nil {
		return err
	}

	return p.DB.Save(&product).Error
}

func (p *Product) Delete(id string) error {
	product, err := p.FindById(id)

	if err != nil {
		return err
	}

	return p.DB.Delete(product).Error
}
