package order

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	sq "github.com/Masterminds/squirrel"
)

func (r *Repository) ListOrders(ctx context.Context, in *dto.ListOrdersIn) (*dto.ListOrderOut, error) {
	var response dto.ListOrderOut
	query, args, err := r.sq.
		Select("*").
		From("orders").
		Where(sq.Eq{"id": in.Ids}).
		Limit(in.Limit).
		Offset(in.Offset).
		ToSql()

	if err != nil {
		return nil, err
	}

	var dbItems []dbOrder

	err = r.db.SelectContext(ctx, &dbItems, query, args...)
	if err != nil {
		return nil, err
	}

	response.Items = listToDto(dbItems)
	response.Count = uint64(len(dbItems))

	return &response, nil
}
