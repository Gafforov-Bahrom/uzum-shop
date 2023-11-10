package main

import (
	"github.com/Gafforov-Bahrom/uzum_shop/internal/service/basket"
	"github.com/Gafforov-Bahrom/uzum_shop/internal/service/order"
	"github.com/Gafforov-Bahrom/uzum_shop/internal/service/product"
	"github.com/Gafforov-Bahrom/uzum_shop/internal/service/user"
)

func (sp *serviceProvider) GetProductService() *product.Service {
	if sp.serviceProducts == nil {
		sp.serviceProducts = product.NewService(
			sp.GetProductRepository(),
		)
	}
	return sp.serviceProducts
}

func (sp *serviceProvider) GetBasketService() *basket.Service {
	if sp.serviceBaskets == nil {
		sp.serviceBaskets = basket.NewService(
			sp.GetBasketRepository(),
		)
	}
	return sp.serviceBaskets
}

func (sp *serviceProvider) GetOrderService() *order.Service {
	if sp.serviceOrders == nil {
		sp.serviceOrders = order.NewService(
			sp.GetOrderRepository(),
			// sp.GetProductRepository(),
			// sp.GetBasketRepository(),
			sp.GetUserRepository(),
		)
	}
	return sp.serviceOrders
}

func (sp *serviceProvider) GetUserService() *user.Service {
	if sp.serviceUsers == nil {
		sp.serviceUsers = user.NewService(
			sp.GetUserRepository(),
		)
	}
	return sp.serviceUsers
}
