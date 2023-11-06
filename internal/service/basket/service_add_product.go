package basket

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (s *Service) AddProduct(ctx context.Context, in *dto.Basket) (dto.TypeID, error) {
	return s.repo.CreateBasket(ctx, in)
}
