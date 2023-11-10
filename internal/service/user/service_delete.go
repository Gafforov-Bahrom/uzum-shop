package user

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (s *Service) DeleteUser(ctx context.Context, in dto.TypeID) error {
	return s.repo.DeleteUser(ctx, in)
}
