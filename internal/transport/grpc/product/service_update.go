package product

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errUpdateProductNotFound = status.Error(codes.NotFound, "Product not found")

func (s *grpcService) UpdateProduct(ctx context.Context, in *desc.UpdateProductRequest) (*desc.UpdateProductResponse, error) {
	out, err := s.productService.UpdateProduct(ctx, &dto.Product{
		Id:          dto.TypeID(in.Id),
		Name:        in.Name,
		Price:       in.Price,
		Description: in.Description,
		Count:       in.Count,
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errUpdateProductNotFound
		}
		return nil, err
	}

	return &desc.UpdateProductResponse{
		Id: uint64(out),
	}, nil
}
