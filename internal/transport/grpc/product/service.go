package product

import desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"

type grpcService struct {
	desc.UnimplementedShopServiceServer
	productService ProductService
	basketService  BasketService
	orderService   OrderService
}

func NewService(productService ProductService, basketService BasketService, orderService OrderService) desc.ShopServiceServer {
	return &grpcService{
		productService: productService,
		basketService:  basketService,
		orderService:   orderService,
	}
}
