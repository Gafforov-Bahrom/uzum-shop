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

var errGetProductNotFound = status.Error(codes.NotFound, "product not found")

func (s *grpcService) GetProduct(ctx context.Context, in *desc.GetProductRequest) (*desc.Product, error) {
	out, err := s.productService.GetProduct(ctx, dto.TypeID(in.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errGetProductNotFound
		}
		return nil, err
	}
	response := dtoToDesc(out)
	return response, nil
}
