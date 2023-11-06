package product

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var errDeleteProductNotFound = status.Error(codes.NotFound, "Basket not found")

func (s *grpcService) DeleteProduct(ctx context.Context, in *desc.DeleteProductRequest) (*emptypb.Empty, error) {
	err := s.basketService.DeleteProduct(ctx, &dto.Basket{
		UserId:    dto.TypeID(in.UserId),
		ProductId: dto.TypeID(in.ProductId),
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errDeleteProductNotFound
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
