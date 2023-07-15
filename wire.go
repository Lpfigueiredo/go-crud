//go:build wireinject
// +build wireinject

package main

import (
	"lpfigueiredo/go-crud/product"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func initProductAPI(db *gorm.DB) product.ProductAPI {
	wire.Build(product.ProductRepositorySet, product.ProductServiceSet, product.ProductAPISet)

	return product.ProductAPI{}
}
