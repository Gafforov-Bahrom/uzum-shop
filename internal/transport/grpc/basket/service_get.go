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

var errGetBasketNotFound = status.Error(codes.NotFound, "basket not found")

func (s *grpcService) GetBasket(ctx context.Context, in *desc.GetBasketsRequest) (*desc.GetBasketsResponse, error) {
	userId, err := s.getUserId(ctx, in.AccessToken)
	if err != nil {
		return nil, err
	}

	out, err := s.basketService.ListBaskets(ctx, dto.TypeID(userId))

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
