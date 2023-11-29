package order

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errListOrdersNotFount = status.Error(codes.NotFound, "order not found")

func (s *grpcService) ListOrders(ctx context.Context, in *desc.ListOrdersRequest) (*desc.ListOrdersResponse, error) {
	dtoList := make([]dto.TypeID, len(in.Ids))
	for i, v := range in.Ids {
		dtoList[i] = dto.TypeID(v)
	}
	out, err := s.orderService.ListOrders(ctx, &dto.ListOrdersIn{
		Limit:  in.Limit,
		Offset: in.Offset,
		Ids:    dtoList,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errListOrdersNotFount
		}
		return nil, err
	}

	return &desc.ListOrdersResponse{
		Count: out.Count,
		Items: listToDesc(out.Items),
	}, nil
}
