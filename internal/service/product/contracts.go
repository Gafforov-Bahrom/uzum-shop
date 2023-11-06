package product

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

type ProductRepository interface {
	List(context.Context, *dto.ListProductIn) (*dto.ListProductOut, error)
	GetProduct(context.Context, dto.TypeID) (*dto.Product, error)
}
