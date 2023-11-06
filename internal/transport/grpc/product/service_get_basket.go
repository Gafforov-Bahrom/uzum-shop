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

var errGetBasketNotFound = status.Error(codes.NotFound, "basket not found")

func (s *grpcService) GetBasket(ctx context.Context, in *desc.GetBasketsRequest) (*desc.GetBasketsResponse, error) {
	out, err := s.basketService.ListBaskets(ctx, dto.TypeID(in.UserId))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errGetBasketNotFound
		}
		return nil, err
	}

	return &desc.GetBasketsResponse{
		Items: listToDescBasket(out),
	}, nil
}
