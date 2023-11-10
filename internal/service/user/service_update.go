package user

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (s *Service) UpdateUser(ctx context.Context, in *dto.User) (*dto.User, error) {
	return s.repo.UpdateUser(ctx, in)
}
