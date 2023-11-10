package order

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (s *Service) CreateOrder(ctx context.Context, in dto.TypeID) (dto.TypeID, error) {
	var response dto.Order
	listDto := make([]dto.TypeID, 1)
	listDto[0] = in
	user, err := s.userRepo.ListUsers(ctx, listDto)
	if err != nil {
		return 0, err
	}
	response.UserId = user[0].Id
	response.Address = user[0].Address
	response.CoordinateAddressX = user[0].CoordinateAddressX
	response.CoordinateAddressY = user[0].CoordinateAddressY

	return s.repo.CreateOrder(ctx, &response)
}
