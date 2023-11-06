package basket

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (s *Service) UpdateBasket(ctx context.Context, in *dto.Basket) (*dto.Basket, error) {
	return s.repo.UpdateBasket(ctx, in)
}
