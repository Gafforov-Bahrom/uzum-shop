package main

import (
	"github.com/Gafforov-Bahrom/uzum_shop/internal/repository/basket"
	"github.com/Gafforov-Bahrom/uzum_shop/internal/repository/product"
)

func (sp *serviceProvider) GetProductRepository() *product.Repository {
	return product.NewRepository(
		sp.GetDB(),
	)
}

func (sp *serviceProvider) GetBasketRepository() *basket.Repository {
	return basket.NewRepository(
		sp.GetDB(),
	)
}
