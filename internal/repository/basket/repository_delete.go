package basket

import (
	"context"
	"database/sql"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (r *Repository) DeleteProduct(ctx context.Context, in *dto.Basket) error {
	query, args, err := r.sq.
		Delete("baskets").
		Where("user_id=?", in.UserId).
		Where("product_id=?", in.ProductId).
		ToSql()

	if err != nil {
		return err
	}

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	count, _ := result.RowsAffected()
	if count == 0 {
		return sql.ErrNoRows
	}

	return nil
}
