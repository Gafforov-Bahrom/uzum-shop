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

var errListProductsNotFound = status.Error(codes.NotFound, "product not found")

func (s *grpcService) ListProducts(ctx context.Context, in *desc.ListProductsRequest) (*desc.ListProductsResponse, error) {
	out, err := s.productService.List(ctx, &dto.ListProductIn{
		Limit:  in.Limit,
		Offset: in.Offset,
		Query:  in.Query,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errListProductsNotFound
		}
		return nil, err
	}

	return &desc.ListProductsResponse{
		Count: out.Count,
		Items: listToDesc(out.Items),
	}, nil
}
