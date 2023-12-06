package order

import (
	"context"
	"time"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	"github.com/Masterminds/squirrel"
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

	err = r.updateProducts(ctx, tx, response)
	if err != nil {
		return 0, err
	}

	err = r.moveBaskets(ctx, tx, response)
	if err != nil {
		return 0, nil
	}

	return response, nil
}

func (r *Repository) updateProducts(ctx context.Context, tx *sqlx.Tx, orderId dto.TypeID) error {

	query, args, err := r.sq.Update("products").
		Prefix(`WITH temp AS (
			SELECT product_id, count FROM baskets WHERE id IN (
				SELECT unnest(items) FROM orders WHERE id = ?
			)
		)`, orderId).
		Set("count", squirrel.Expr(`count - (SELECT count FROM temp WHERE product_id=id)`)).
		Where("id IN (SELECT product_id FROM temp)").
		ToSql()

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)

	return err
}

func (r *Repository) moveBaskets(ctx context.Context, tx *sqlx.Tx, orderId dto.TypeID) error {
	query := `
		INSERT INTO order_items(id, product_id, count, order_id)
		SELECT id, product_id, count, $1 FROM baskets WHERE id IN (SELECT unnest(items) FROM orders WHERE id = $2)
	`

	_, err := tx.ExecContext(ctx, query, orderId, orderId)
	if err != nil {
		return err
	}

	query = `DELETE FROM baskets WHERE id IN (SELECT unnest(items) FROM orders WHERE id = $1)`

	_, err = tx.ExecContext(ctx, query, orderId)
	return err
}
