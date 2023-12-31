package order

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

var errDeleteOrderNotFound = status.Error(codes.NotFound, "Order not found")

func (s *grpcService) DeleteOrder(ctx context.Context, in *desc.DeleteOrderRequest) (*emptypb.Empty, error) {
	userId, err := s.getUserId(ctx, in.AccessToken)
	if err != nil {
		return nil, err
	}

	err = s.orderService.DeleteOrder(ctx, dto.DeleteOrderRequest{
		UserId: dto.TypeID(userId),
		Id:     dto.TypeID(in.Id),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &emptypb.Empty{}, errDeleteOrderNotFound
		}
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
