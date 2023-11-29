package order

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

type OrderRepository interface {
	CreateOrder(context.Context, *dto.Order) (dto.TypeID, error)
	DeleteOrder(context.Context, dto.TypeID) error
	ListOrders(context.Context, *dto.ListOrdersIn) (*dto.ListOrderOut, error)
}

type UserRepository interface {
	ListUsers(context.Context, []dto.TypeID) ([]*dto.User, error)
}

type BasketRepository interface {
	ListBaskets(context.Context, dto.TypeID) ([]*dto.Basket, error)
}
