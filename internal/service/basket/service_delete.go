package basket

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (s *Service) DeleteProduct(ctx context.Context, in *dto.Basket) error {
	return s.repo.DeleteProduct(ctx, in)
}
