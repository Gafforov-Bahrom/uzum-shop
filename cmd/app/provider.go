package main

import (
	"github.com/Gafforov-Bahrom/uzum_shop/internal/config"
	repoBaskets "github.com/Gafforov-Bahrom/uzum_shop/internal/repository/basket"
	repoOrders "github.com/Gafforov-Bahrom/uzum_shop/internal/repository/order"
	repoProduct "github.com/Gafforov-Bahrom/uzum_shop/internal/repository/product"
	repoUsers "github.com/Gafforov-Bahrom/uzum_shop/internal/repository/user"
	serviceBaskets "github.com/Gafforov-Bahrom/uzum_shop/internal/service/basket"
	serviceOrders "github.com/Gafforov-Bahrom/uzum_shop/internal/service/order"
	serviceProducts "github.com/Gafforov-Bahrom/uzum_shop/internal/service/product"
	serviceUsers "github.com/Gafforov-Bahrom/uzum_shop/internal/service/user"

	"github.com/jmoiron/sqlx"
)

type serviceProvider struct {
	cfg             *config.Config
	db              *sqlx.DB
	repoProducts    *repoProduct.Repository
	repoBaskets     *repoBaskets.Repository
	repoOrders      *repoOrders.Repository
	serviceProducts *serviceProducts.Service
	serviceBaskets  *serviceBaskets.Service
	serviceOrders   *serviceOrders.Service
	repoUsers       *repoUsers.Repository
	serviceUsers    *serviceUsers.Service
}

func NewServiceProvider(cfg *config.Config) *serviceProvider {
	return &serviceProvider{cfg: cfg}
}
