package main

import (
	"github.com/Gafforov-Bahrom/uzum_shop/internal/transport/grpc/product"
	"github.com/Gafforov-Bahrom/uzum_shop/internal/transport/grpc/user"

	desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"
)

func (sp *serviceProvider) GetProductServer() desc.ShopServiceServer {
	return product.NewService(
		sp.GetProductService(),
		sp.GetBasketService(),
		sp.GetOrderService(),
	)
}

func (sp *serviceProvider) GetUserServer() desc.UserServiceServer {
	return user.NewService(
		sp.GetUserService(),
	)
}
