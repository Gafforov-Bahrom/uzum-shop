package user

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (s *Service) ListUsers(ctx context.Context, in []dto.TypeID) ([]*dto.User, error) {
	return s.repo.ListUsers(ctx, in)
}
