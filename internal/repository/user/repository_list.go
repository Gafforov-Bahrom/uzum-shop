package user

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	sq "github.com/Masterminds/squirrel"
)

func (r *Repository) ListUsers(ctx context.Context, in []dto.TypeID) ([]*dto.User, error) {
	query, args, err := r.sq.
		Select("id", "name", "surname", "phone", "login", "password_hash", "role", "address", "coordinate_address_x", "coordinate_address_y").
		From("users").
		Where(sq.Eq{"id": in}).
		ToSql()

	if err != nil {
		return nil, err
	}

	var dbItems []dbUsers

	err = r.db.SelectContext(ctx, &dbItems, query, args...)
	if err != nil {
		return nil, err
	}

	return listToDto(dbItems), nil
}
