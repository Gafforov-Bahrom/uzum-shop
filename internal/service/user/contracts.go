package user

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

type UserRepository interface {
	ListUsers(context.Context, []dto.TypeID) ([]*dto.User, error)
	AddUser(context.Context, dto.User) (dto.TypeID, error)
	UpdateUser(context.Context, *dto.User) (*dto.User, error)
	DeleteUser(context.Context, dto.TypeID) error
}
