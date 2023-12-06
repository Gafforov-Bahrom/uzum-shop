package order

import (
	"context"
	"database/sql"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

func (r *Repository) DeleteOrder(ctx context.Context, in dto.DeleteOrderRequest) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	err = r.cancelProducts(ctx, tx, in.Id)
	if err != nil {
		return err
	}

	query, args, err := r.sq.
		Delete("orders").
		Where("id=?", in.Id).
		Where("user_id=?", in.UserId).
		ToSql()

	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	count, _ := result.RowsAffected()
	if count == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *Repository) cancelProducts(ctx context.Context, tx *sqlx.Tx, orderId dto.TypeID) error {

	// query := `
	// 	WITH temp AS (
	// 		SELECT product_id, count FROM baskets WHERE id IN (
	// 			SELECT unnest(items) FROM orders WHERE id = ?
	// 		)
	// 	)

	// 	UPDATE products SET count = count + (SELECT count FROM temp WHERE product_id=id)
	// `

	query, args, err := r.sq.Update("products").
		Prefix(`WITH temp AS (
			SELECT product_id, count FROM order_items WHERE order_id = ?
		)`, orderId).
		Set("count", squirrel.Expr(`count + (SELECT count FROM temp WHERE product_id=id)`)).
		Where("id IN (SELECT product_id FROM temp)").
		ToSql()

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)

	return err
}
