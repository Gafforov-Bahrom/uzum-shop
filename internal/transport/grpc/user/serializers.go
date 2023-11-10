package user

import (
	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"
)

func dtoToDesc(in *dto.User) *desc.User {
	return &desc.User{
		Id:                 uint64(in.Id),
		Name:               in.Name,
		Surname:            in.Surname,
		Phone:              in.Phone,
		Login:              in.Login,
		PasswordHash:       in.PasswordHash,
		Role:               in.Role,
		Address:            in.Address,
		CoordinateAddressX: float32(in.CoordinateAddressX),
		CoordinateAddressY: float32(in.CoordinateAddressY),
	}
}

func listToDesc(in []*dto.User) []*desc.User {
	out := make([]*desc.User, len(in))
	for i, item := range in {
		out[i] = dtoToDesc(item)
	}
	return out
}
