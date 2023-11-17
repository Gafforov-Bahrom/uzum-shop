package product

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (r *Repository) UpdateProduct(ctx context.Context, in *dto.Product) (dto.TypeID, error) {
	query, args, err := r.sq.
		Update("products").
		Set("name", in.Name).
		Set("price", in.Price).
		Set("description", in.Description).
		Set("count", in.Count).
		Where("id=?", in.Id).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, err
	}

	var dbItem dbProduct

	err = r.db.GetContext(ctx, &dbItem, query, args...)
	if err != nil {
		return 0, err
	}

	return dto.TypeID(dbItem.Id), nil
}
