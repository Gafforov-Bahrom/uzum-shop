package basket

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (r *Repository) UpdateBasket(ctx context.Context, in *dto.Basket) (*dto.Basket, error) {

	query, args, err := r.sq.
		Update("baskets").
		Set("count", in.Count).
		Where("user_id=?", in.UserId).
		Where("product_id=?", in.ProductId).
		Suffix("RETURNING *").
		ToSql()

	if err != nil {
		return nil, err
	}

	var dbItem dbBasket

	err = r.db.GetContext(ctx, &dbItem, query, args...)
	if err != nil {
		return nil, err
	}

	return dbItem.toDTO(), nil
}
