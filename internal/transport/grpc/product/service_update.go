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

var errUpdateProductNotFound = status.Error(codes.NotFound, "Basket not found")

func (s *grpcService) UpdateBasket(ctx context.Context, in *desc.UpdateBasketRequest) (*desc.Basket, error) {
	out, err := s.basketService.UpdateBasket(ctx, &dto.Basket{
		UserId:    dto.TypeID(in.UserId),
		ProductId: dto.TypeID(in.ProductId),
		Count:     in.Count,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errUpdateProductNotFound
		}
		return nil, err
	}

	return dtoToDescBasket(out), nil
}
