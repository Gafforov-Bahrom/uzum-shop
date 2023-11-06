package main

import (
	"github.com/Gafforov-Bahrom/uzum_shop/internal/service/basket"
	"github.com/Gafforov-Bahrom/uzum_shop/internal/service/product"
)

func (sp *serviceProvider) GetProductService() *product.Service {
	return product.NewService(
		sp.GetProductRepository(),
	)
}

func (sp *serviceProvider) GetBasketService() *basket.Service {
	return basket.NewService(
		sp.GetBasketRepository(),
	)
}
