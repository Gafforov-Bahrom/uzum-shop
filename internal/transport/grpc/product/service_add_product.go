package product

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"
)

func (s *grpcService) AddProduct(ctx context.Context, in *desc.AddProductRequest) (*desc.AddProductResponse, error) {
	out, err := s.basketService.AddProduct(ctx, &dto.Basket{
		UserId:    dto.TypeID(in.UserId),
		ProductId: dto.TypeID(in.ProductId),
		Count:     in.Count,
	})

	if err != nil {
		return nil, err
	}

	response := &desc.AddProductResponse{
		Id: uint64(out),
	}

	return response, nil
}
