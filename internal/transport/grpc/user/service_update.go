package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errUpdateUserNotFound = status.Error(codes.NotFound, "User not found")

func (s *grpcService) UpdateUser(ctx context.Context, in *desc.UpdateUserRequest) (*desc.User, error) {
	out, err := s.userService.UpdateUser(ctx, &dto.User{
		Id:                 dto.TypeID(in.Id),
		Name:               in.Name,
		Surname:            in.Surname,
		Phone:              in.Phone,
		Login:              in.Login,
		PasswordHash:       in.PasswordHash,
		Role:               in.Role,
		Address:            in.Address,
		CoordinateAddressX: float64(in.CoordinateAddressX),
		CoordinateAddressY: float64(in.CoordinateAddressY),
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errUpdateUserNotFound
		}
		return nil, err
	}

	return dtoToDesc(out), nil
}
