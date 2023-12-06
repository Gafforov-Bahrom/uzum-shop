package basket

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

var errDeleteBasketProductNotFound = status.Error(codes.NotFound, "Basket not found")

func (s *grpcService) DeleteBasketProduct(ctx context.Context, in *desc.DeleteBasketProductRequest) (*emptypb.Empty, error) {
	// s.loginClient.GetUserId(ctx, &login_v1.GetUserIdRequest{
	// 	A
	// })
	userId, err := s.getUserId(ctx, in.AccessToken)
	if err != nil {
		return nil, err
	}

	err = s.basketService.DeleteProduct(ctx, &dto.Basket{
		UserId:    dto.TypeID(userId),
		ProductId: dto.TypeID(in.ProductId),
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errDeleteBasketProductNotFound
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
