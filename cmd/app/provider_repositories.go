package main

import (
	"github.com/Gafforov-Bahrom/uzum_shop/internal/repository/basket"
	"github.com/Gafforov-Bahrom/uzum_shop/internal/repository/order"
	"github.com/Gafforov-Bahrom/uzum_shop/internal/repository/product"
	"github.com/Gafforov-Bahrom/uzum_shop/internal/repository/user"
)

func (sp *serviceProvider) GetProductRepository() *product.Repository {
	if sp.repoProducts == nil {
		sp.repoProducts = product.NewRepository(
			sp.GetDB(),
		)
	}

	return sp.repoProducts
}

func (sp *serviceProvider) GetBasketRepository() *basket.Repository {
	if sp.repoBaskets == nil {
		sp.repoBaskets = basket.NewRepository(
			sp.GetDB(),
		)
	}
	return sp.repoBaskets
}

func (sp *serviceProvider) GetOrderRepository() *order.Repository {
	if sp.repoOrders == nil {
		sp.repoOrders = order.NewRepository(
			sp.GetDB(),
		)
	}
	return sp.repoOrders
}

func (sp *serviceProvider) GetUserRepository() *user.Repository {
	if sp.repoUsers == nil {
		sp.repoUsers = user.NewRepository(
			sp.GetDB(),
		)
	}
	return sp.repoUsers
}
