package basket

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
)

func (r *Repository) CreateBasket(ctx context.Context, in *dto.Basket) (dto.TypeID, error) {
	var response dto.TypeID
	query, args, err := r.sq.
		Insert("baskets").
		Columns(
			"user_id",
			"product_id",
			"count",
		).
		Values(
			in.UserId,
			in.ProductId,
			in.Count,
		).
		Suffix("ON CONFLICT(user_id , product_id ) DO UPDATE SET count=baskets.count+excluded.count").
		//Suffix("ON CONFLICT ON CONSTRAINT document_numbers_prefix_contractor_id_key DO UPDATE SET number=document_numbers.number+1").
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}

	var dbItem dbBasket

	err = r.db.GetContext(ctx, &dbItem, query, args...)
	if err != nil {
		return 0, err
	}

	response = dto.TypeID(dbItem.Id)

	return response, nil
}
