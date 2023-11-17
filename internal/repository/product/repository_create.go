package product

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (r *Repository) CreateProduct(ctx context.Context, in dto.Product) (dto.TypeID, error) {
	query, args, err := r.sq.
		Insert("products").
		Columns(
			"name",
			"price",
			"description",
			"count",
		).
		Values(
			in.Name,
			in.Price,
			in.Description,
			in.Count,
		).
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
