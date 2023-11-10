package user

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"
)

func (s *grpcService) AddUser(ctx context.Context, in *desc.AddUserRequest) (*desc.AddUserResponse, error) {
	dtoUser := func(in *desc.AddUserRequest) dto.User {
		return dto.User{
			Name:               in.Name,
			Surname:            in.Surname,
			Phone:              in.Phone,
			Login:              in.Login,
			PasswordHash:       in.PasswordHash,
			Role:               in.Role,
			Address:            in.Address,
			CoordinateAddressX: float64(in.CoordinateAddressX),
			CoordinateAddressY: float64(in.CoordinateAddressY),
		}
	}(in)

	out, err := s.userService.AddUser(ctx, dtoUser)
	if err != nil {
		return nil, err
	}

	return &desc.AddUserResponse{
		Id: uint64(out),
	}, nil

}
