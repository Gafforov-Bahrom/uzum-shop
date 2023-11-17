package product

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"
)

func (s *grpcService) CreateProduct(ctx context.Context, in *desc.CreateProductRequest) (*desc.CreateProductResponse, error) {
	dtoProduct := func(in *desc.CreateProductRequest) dto.Product {
		return dto.Product{
			Name:        in.Name,
			Description: in.Description,
			Count:       in.Count,
			Price:       in.Price,
		}
	}(in)

	out, err := s.productService.CreateProduct(ctx, dtoProduct)
	if err != nil {
		return nil, err
	}

	return &desc.CreateProductResponse{
		Id: uint64(out),
	}, nil
}
