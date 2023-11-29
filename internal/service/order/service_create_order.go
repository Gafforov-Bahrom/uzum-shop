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

	items, err := s.basketRepo.ListBaskets(ctx, response.UserId)
	if err != nil {
		return 0, err
	}

	temp := make([]uint64, len(items))

	for i, item := range items {
		temp[i] = uint64(item.Id)
	}
	response.Items = temp

	return s.repo.CreateOrder(ctx, &response)
}
