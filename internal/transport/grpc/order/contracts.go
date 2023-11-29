package order

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

type OrderService interface {
	CreateOrder(context.Context, dto.TypeID) (dto.TypeID, error)
	DeleteOrder(context.Context, dto.TypeID) error
	ListOrders(context.Context, *dto.ListOrdersIn) (*dto.ListOrderOut, error)
}
