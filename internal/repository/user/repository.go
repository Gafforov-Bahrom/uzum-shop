package user

import (
	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
	sq sq.StatementBuilderType
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

type dbUsers struct {
	Id                 dto.TypeID `db:"id"`
	Name               string     `db:"name"`
	Surname            string     `db:"surname"`
	Phone              string     `db:"phone"`
	Login              string     `db:"login"`
	PasswordHash       string     `db:"password_hash"`
	Role               string     `db:"role"`
	Address            string     `db:"address"`
	CoordinateAddressX float64    `db:"coordinate_address_x"`
	CoordinateAddressY float64    `db:"coordinate_address_y"`
}

func (item dbUsers) toDTO() *dto.User {
	return &dto.User{
		Id:                 dto.TypeID(item.Id),
		Name:               item.Name,
		Surname:            item.Surname,
		Phone:              item.Phone,
		Login:              item.Login,
		PasswordHash:       item.PasswordHash,
		Role:               item.Role,
		Address:            item.Address,
		CoordinateAddressX: item.CoordinateAddressX,
		CoordinateAddressY: item.CoordinateAddressY,
	}
}

func listToDto(items []dbUsers) []*dto.User {
	response := make([]*dto.User, len(items))
	for i, item := range items {
		response[i] = item.toDTO()
	}
	return response
}
