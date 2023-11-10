package user

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (r *Repository) UpdateUser(ctx context.Context, in *dto.User) (*dto.User, error) {
	query, args, err := r.sq.
		Update("users").
		Set("name", in.Name).
		Set("surname", in.Surname).
		Set("phone", in.Phone).
		Set("login", in.Login).
		Set("password_hash", in.PasswordHash).
		Set("role", in.Role).
		Set("address", in.Address).
		Set("coordinate_address_x", in.CoordinateAddressX).
		Set("coordinate_address_y", in.CoordinateAddressY).
		Where("id=?", in.Id).
		Suffix("RETURNING *").
		ToSql()

	if err != nil {
		return nil, err
	}

	var dbItem dbUsers

	err = r.db.GetContext(ctx, &dbItem, query, args...)
	if err != nil {
		return nil, err
	}

	return dbItem.toDTO(), nil
}
