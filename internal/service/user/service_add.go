package user

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (s *Service) AddUser(ctx context.Context, in dto.User) (dto.TypeID, error) {
	return s.repo.AddUser(ctx, in)
}
