package main

import (
	"github.com/Gafforov-Bahrom/uzum_shop/internal/transport/grpc/product"
	desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"
)

func (sp *serviceProvider) GetProductServer() desc.ShopServiceServer {
	return product.NewService(
		sp.GetProductService(),
		sp.GetBasketService(),
	)
}
