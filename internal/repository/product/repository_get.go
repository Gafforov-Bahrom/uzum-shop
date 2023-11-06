package product

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (r *Repository) GetProduct(ctx context.Context, id dto.TypeID) (*dto.Product, error) {
	var response dto.Product
	query, args, err := r.sq.
		Select("id", "name", "description", "price", "count").
		From("products").
		Where("id=?", id).
		ToSql()
	if err != nil {
		return nil, err
	}
	var dbItem dbProduct

	err = r.db.GetContext(ctx, &dbItem, query, args...)
	if err != nil {
		return nil, err
	}
	response = func(dbItem dbProduct) dto.Product {
		return dto.Product{
			Id:          dto.TypeID(dbItem.Id),
			Name:        dbItem.Name,
			Description: dbItem.Description,
			Price:       dbItem.Price,
			Count:       dbItem.Count,
		}
	}(dbItem)

	return &response, nil
}
