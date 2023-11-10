package product

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

type ProductService interface {
	List(context.Context, *dto.ListProductIn) (*dto.ListProductOut, error)
	GetProduct(context.Context, dto.TypeID) (*dto.Product, error)
}

type BasketService interface {
	AddProduct(context.Context, *dto.Basket) (dto.TypeID, error)
	UpdateBasket(context.Context, *dto.Basket) (*dto.Basket, error)
	DeleteProduct(context.Context, *dto.Basket) error
	ListBaskets(context.Context, dto.TypeID) ([]*dto.Basket, error)
}

type OrderService interface {
	CreateOrder(context.Context, dto.TypeID) (dto.TypeID, error)
	DeleteOrder(context.Context, dto.TypeID) error
}
