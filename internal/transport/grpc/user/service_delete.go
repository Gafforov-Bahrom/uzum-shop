package user

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

var errDeleteUserNotFound = status.Error(codes.NotFound, "User not found")

func (s *grpcService) DeleteUser(ctx context.Context, in *desc.DeleteUserRequest) (*emptypb.Empty, error) {
	err := s.userService.DeleteUser(ctx, dto.TypeID(in.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errDeleteUserNotFound
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
