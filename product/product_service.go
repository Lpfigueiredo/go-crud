package product

import "github.com/google/wire"

type ProductService struct {
	ProductRepository ProductRepository
}

var ProductServiceSet = wire.NewSet(wire.Struct(new(ProductService), "*"))

func (p *ProductService) FindAll() []Product {
	return p.ProductRepository.FindAll()
}

func (p *ProductService) FindByID(id uint) Product {
	return p.ProductRepository.FindByID(id)
}

func (p *ProductService) Save(product Product) Product {
	p.ProductRepository.Save(product)

	return product
}

func (p *ProductService) Delete(product Product) {
	p.ProductRepository.Delete(product)
}
