package order

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (s *Service) ListOrders(ctx context.Context, in *dto.ListOrdersIn) (*dto.ListOrderOut, error) {
	return s.repo.ListOrders(ctx, in)
}
