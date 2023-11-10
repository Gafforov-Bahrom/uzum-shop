package user

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (r *Repository) AddUser(ctx context.Context, in dto.User) (dto.TypeID, error) {
	query, args, err := r.sq.
		Insert("users").
		Columns(
			"name",
			"surname",
			"address",
			"phone",
			"login",
			"password_hash",
			"role",
			"coordinate_address_x",
			"coordinate_address_y",
		).
		Values(
			in.Name,
			in.Surname,
			in.Address,
			in.Phone,
			in.Login,
			in.PasswordHash,
			in.Role,
			in.CoordinateAddressX,
			in.CoordinateAddressY,
		).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, err
	}

	var dbItem dbUsers

	err = r.db.GetContext(ctx, &dbItem, query, args...)
	if err != nil {
		return 0, err
	}

	return dto.TypeID(dbItem.Id), nil
}
