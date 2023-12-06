package order

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"
)

func (s *grpcService) CreateOrder(ctx context.Context, in *desc.CreateOrderRequest) (*desc.CreateOrderResponse, error) {
	userId, err := s.getUserId(ctx, in.AccessToken)
	if err != nil {
		return nil, err
	}

	out, err := s.orderService.CreateOrder(ctx, dto.TypeID(userId))
	if err != nil {
		return nil, err
	}

	return &desc.CreateOrderResponse{
		Id: uint64(out),
	}, nil
}
