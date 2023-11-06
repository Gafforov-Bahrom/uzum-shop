package basket

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (r *Repository) ListBaskets(ctx context.Context, userId dto.TypeID) ([]*dto.Basket, error) {
	query, args, err := r.sq.
		Select("id", "user_id", "product_id", "count").
		From("baskets").
		Where("user_id=?", userId).
		ToSql()
	if err != nil {
		return nil, err
	}

	var dbItems []dbBasket

	err = r.db.SelectContext(ctx, &dbItems, query, args...)

	if err != nil {
		return nil, err
	}

	return ListToDto(dbItems), nil
}
