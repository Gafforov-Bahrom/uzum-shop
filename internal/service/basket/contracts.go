package basket

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

type BasketRepository interface {
	CreateBasket(context.Context, *dto.Basket) (dto.TypeID, error)
	UpdateBasket(context.Context, *dto.Basket) (*dto.Basket, error)
	DeleteProduct(context.Context, *dto.Basket) error
	ListBaskets(context.Context, dto.TypeID) ([]*dto.Basket, error)
}
