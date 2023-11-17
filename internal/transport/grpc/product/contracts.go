package product

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

type ProductService interface {
	List(context.Context, *dto.ListProductIn) (*dto.ListProductOut, error)
	GetProduct(context.Context, dto.TypeID) (*dto.Product, error)
	CreateProduct(context.Context, dto.Product) (dto.TypeID, error)
	UpdateProduct(context.Context, *dto.Product) (dto.TypeID, error)
	DeleteProduct(context.Context, dto.TypeID) error
}
