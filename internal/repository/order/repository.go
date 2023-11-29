package order

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Repository struct {
	db *sqlx.DB
	sq sq.StatementBuilderType
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

type dbOrder struct {
	Id                 dto.TypeID    `db:"id"`
	UserId             dto.TypeID    `db:"user_id"`
	Address            string        `db:"address"`
	CoordinateAddressX float64       `db:"coordinate_address_x"`
	CoordinateAddressY float64       `db:"coordinate_address_y"`
	CoordinatePointX   float64       `db:"coordinate_point_x"`
	CoordinatePointY   float64       `db:"coordinate_point_y"`
	CreatedAt          time.Time     `db:"create_at"`
	StartedAt          sql.NullTime  `db:"start_at"`
	DeliveryAt         sql.NullTime  `db:"delivery_at"`
	CourierId          sql.NullInt64 `db:"courier_id"`
	DeliveryStatus     int64         `db:"delivery_status"`
	Items              pq.Int64Array `db:"items"`
}

func (item dbOrder) toDTO() *dto.Order {
	dtoItem := make([]uint64, len(item.Items))
	for i, v := range item.Items {
		dtoItem[i] = uint64(v)
	}
	fmt.Println(dtoItem, len(item.Items), item.Items)
	return &dto.Order{
		Id:                 item.Id,
		UserId:             item.UserId,
		Address:            item.Address,
		CoordinateAddressX: item.CoordinateAddressX,
		CoordinateAddressY: item.CoordinateAddressY,
		CoordinatePointX:   item.CoordinatePointX,
		CoordinatePointY:   item.CoordinatePointY,
		CreatedAt:          item.CreatedAt,
		StartedAt:          item.StartedAt.Time,
		DeliveryAt:         item.DeliveryAt.Time,
		CourierId:          dto.TypeID(item.CourierId.Int64),
		DeliveryStatus:     item.DeliveryStatus,
		Items:              dtoItem,
	}
}

func listToDto(items []dbOrder) []*dto.Order {
	response := make([]*dto.Order, len(items))
	for i, item := range items {
		response[i] = item.toDTO()
	}
	return response
}
