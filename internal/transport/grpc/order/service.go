package order

import desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"

type grpcService struct {
	desc.UnimplementedOrderServiceServer
	orderService OrderService
}

func NewService(orderService OrderService) desc.OrderServiceServer {
	return &grpcService{
		orderService: orderService,
	}
}
