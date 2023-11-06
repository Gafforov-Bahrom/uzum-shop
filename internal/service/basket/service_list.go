package basket

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (s *Service) ListBaskets(ctx context.Context, UserId dto.TypeID) ([]*dto.Basket, error) {
	return s.repo.ListBaskets(ctx, UserId)
}
