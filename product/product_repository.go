package product

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

var ProductRepositorySet = wire.NewSet(wire.Struct(new(ProductRepository), "*"))

func (p *ProductRepository) FindAll() []Product {
	var products []Product
	p.DB.Find(&products)

	return products
}

func (p *ProductRepository) FindByID(id uint) Product {
	var product Product
	p.DB.First(&product, id)

	return product
}

func (p *ProductRepository) Save(product Product) Product {
	p.DB.Save(&product)

	return product
}

func (p *ProductRepository) Delete(product Product) {
	p.DB.Delete(&product)
}
