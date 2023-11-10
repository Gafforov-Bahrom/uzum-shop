package order

import (
	"context"
	"time"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (r *Repository) CreateOrder(ctx context.Context, in *dto.Order) (dto.TypeID, error) {
	var response dto.TypeID
	query, args, err := r.sq.
		Insert("orders").
		Columns(
			"user_id",
			"address",
			"coordinate_address_x",
			"coordinate_address_y",
			"coordinate_point_x",
			"coordinate_point_y",
			"create_at",
			"start_at",
			"delivery_at",
			"courier_id",
			"delivery_status",
		).
		Values(
			in.UserId,
			in.Address,
			in.CoordinateAddressX,
			in.CoordinateAddressY,
			in.CoordinatePointX,
			in.CoordinatePointY,
			time.Now(),
			nil,
			nil,
			nil,
			nil,
		).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, err
	}

	var dbItem dbOrder

	err = r.db.GetContext(ctx, &dbItem, query, args...)
	if err != nil {
		return 0, err
	}

	response = dto.TypeID(dbItem.Id)

	return response, nil
	// return 0, nil
}
