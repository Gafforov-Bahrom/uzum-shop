package dto

import "time"

type Order struct {
	Id                 TypeID
	UserId             TypeID
	Address            string
	CoordinateAddressX float64
	CoordinateAddressY float64
	CoordinatePointX   float64
	CoordinatePointY   float64
	CreatedAt          time.Time
	StartedAt          time.Time
	DeliveryAt         time.Time
	CourierId          TypeID
	DeliveryStatus     int64
	Items              []uint64
}

type ListOrdersIn struct {
	Limit  uint64
	Offset uint64
	Ids    []TypeID
}

type ListOrderOut struct {
	Count uint64
	Items []*Order
}

type DeleteOrderRequest struct {
	UserId TypeID
	Id     TypeID
}
