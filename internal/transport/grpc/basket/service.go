package basket

import desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"

type grpcService struct {
	desc.UnimplementedBasketServiceServer
	basketService BasketService
}

func NewService(basketService BasketService) desc.BasketServiceServer {
	return &grpcService{
		basketService: basketService,
	}
}
