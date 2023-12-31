package order

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (s *Service) DeleteOrder(ctx context.Context, in dto.DeleteOrderRequest) error {
	return s.repo.DeleteOrder(ctx, in)
}
