package product

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (r *Repository) List(ctx context.Context, in *dto.ListProductIn) (*dto.ListProductOut, error) {
	var response dto.ListProductOut
	query, args, err := r.sq.
		Select("id", "name", "description", "price", "count").
		From("products").
		Where("name ILIKE ?", "%"+in.Query+"%").
		// Where("id IN (?)", in.Query).
		Limit(in.Limit).
		Offset(in.Offset).
		ToSql()
	if err != nil {
		return nil, err
	}

	var dbItems []dbProduct

	err = r.db.SelectContext(ctx, &dbItems, query, args...)
	if err != nil {
		return nil, err
	}
	response.Items = listToDto(dbItems)

	// rows, err := r.db.QueryContext(ctx, query, args...)
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var product dto.Product
	// 	err = rows.Scan(
	// 		&product.Id,
	// 		&product.Name,
	// 	)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	response.Items = append(response.Items, &product)
	// }

	query, args, err = r.sq.Select("count(id) AS products_count").
		From("products").
		Where("name ILIKE ?", "%"+in.Query+"%").
		ToSql()

	if err != nil {
		return nil, err
	}

	err = r.db.GetContext(ctx, &response.Count, query, args...)

	// err = r.db.QueryRow(query, args...).Scan(&response.Count)
	if err != nil {
		return nil, err
	}
	return &response, nil

	// Bunaqa qilsa ham bo'ladi

	// query := `
	// 	SELECT
	// 		id,
	// 		name
	// 	FROM products
	// 	WHERE name ILIKE $1
	// 	LIMIT $2
	// 	OFFSET $3
	// `
	// r.db.QueryContext(ctx, query,
	// 	"%"+in.Query+"%",
	// 	in.Limit,
	// 	in.Offset,
	// )
}
