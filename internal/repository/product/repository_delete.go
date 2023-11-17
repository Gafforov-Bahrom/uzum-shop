package product

import (
	"context"
	"database/sql"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (r *Repository) DeleteProduct(ctx context.Context, in dto.TypeID) error {
	query, args, err := r.sq.
		Delete("products").
		Where("id=?", in).
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
