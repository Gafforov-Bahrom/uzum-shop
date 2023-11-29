package main

import (
	"github.com/Gafforov-Bahrom/uzum_shop/internal/transport/http"
	"github.com/Gafforov-Bahrom/uzum_shop/internal/transport/http/shop"
	"github.com/gin-gonic/gin"
)

func (sp *serviceProvider) GetShopRouter() *shop.Router {
	return shop.NewRouter(
		sp.GetProductService(),
		sp.GetBasketService(),
		sp.GetOrderService(),
		sp.GetLoginClient(),
	)
}

func (sp *serviceProvider) GetHTTPServer() *http.Server {
	return http.NewServer(sp.cfg.HTTPPort, gin.Default(), sp.GetShopRouter())
}
