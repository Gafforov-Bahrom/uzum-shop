package product

import desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"

type grpcService struct {
	desc.UnimplementedShopServiceServer
	productService ProductService
}

func NewService(productService ProductService) desc.ShopServiceServer {
	return &grpcService{
		productService: productService,
	}
}
