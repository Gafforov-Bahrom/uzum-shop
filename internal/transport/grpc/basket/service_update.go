package basket

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errUpdateBasketProductNotFound = status.Error(codes.NotFound, "Basket not found")

func (s *grpcService) UpdateBasket(ctx context.Context, in *desc.UpdateBasketRequest) (*desc.Basket, error) {
	userId, err := s.getUserId(ctx, in.AccessToken)
	if err != nil {
		return nil, err
	}
	out, err := s.basketService.UpdateBasket(ctx, &dto.Basket{
		UserId:    dto.TypeID(userId),
		ProductId: dto.TypeID(in.ProductId),
		Count:     in.Count,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errUpdateBasketProductNotFound
		}
		return nil, err
	}

	return dtoToDescBasket(out), nil
}
