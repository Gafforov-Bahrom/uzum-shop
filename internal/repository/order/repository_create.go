package order

import (
	"context"
	"time"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func (r *Repository) CreateOrder(ctx context.Context, in *dto.Order) (dto.TypeID, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	dtoItem := make([]int64, len(in.Items))
	for i, v := range in.Items {
		dtoItem[i] = int64(v)
	}
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
			"items",
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
			0,
			pq.Int64Array(dtoItem),
		).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, err
	}

	var dbItem dbOrder

	err = tx.GetContext(ctx, &dbItem, query, args...)
	if err != nil {
		return 0, err
	}

	response = dto.TypeID(dbItem.Id)

	_ = r.updateProducts(tx, response)

	return response, nil
}

func (r *Repository) updateProducts(tx *sqlx.Tx, orderId dto.TypeID) error {
	return nil
}
