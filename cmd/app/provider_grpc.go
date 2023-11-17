package main

import (
	"github.com/Gafforov-Bahrom/uzum_shop/internal/transport/grpc/basket"
	"github.com/Gafforov-Bahrom/uzum_shop/internal/transport/grpc/order"
	"github.com/Gafforov-Bahrom/uzum_shop/internal/transport/grpc/product"
	"github.com/Gafforov-Bahrom/uzum_shop/internal/transport/grpc/user"

	desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"
)

func (sp *serviceProvider) GetProductServer() desc.ShopServiceServer {
	return product.NewService(
		sp.GetProductService(),
	)
}

func (sp *serviceProvider) GetBasketServer() desc.BasketServiceServer {
	return basket.NewService(
		sp.GetBasketService(),
	)
}

func (sp *serviceProvider) GetOrderServer() desc.OrderServiceServer {
	return order.NewService(
		sp.GetOrderService(),
	)
}

func (sp *serviceProvider) GetUserServer() desc.UserServiceServer {
	return user.NewService(
		sp.GetUserService(),
	)
}
