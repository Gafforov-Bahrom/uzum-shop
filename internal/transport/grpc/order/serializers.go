package order

import (
	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func dtoToDesc(in *dto.Order) *desc.Order {
	return &desc.Order{
		Id:                 uint64(in.Id),
		UserId:             uint64(in.UserId),
		Address:            in.Address,
		CoordinateAddressX: float32(in.CoordinateAddressX),
		CoordinateAddressY: float32(in.CoordinateAddressY),
		CoordinatePointX:   float32(in.CoordinatePointX),
		CoordinatePointY:   float32(in.CoordinatePointY),
		CreatedAt:          timestamppb.New(in.CreatedAt),
		StartedAt:          timestamppb.New(in.StartedAt),
		DeliveryAt:         timestamppb.New(in.DeliveryAt),
		CourierId:          uint64(in.CourierId),
		DeliveryStatus:     uint64(in.DeliveryStatus),
		Items:              in.Items,
	}
}

func listToDesc(in []*dto.Order) []*desc.Order {
	out := make([]*desc.Order, len(in))
	for i, item := range in {
		out[i] = dtoToDesc(item)
	}
	return out
}
