package order

import (
	"time"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
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
	Id                 dto.TypeID `db:"id"`
	UserId             dto.TypeID `db:"user_id,"`
	Address            string     `db:"address,"`
	CoordinateAddressX float64    `db:"coordinate_address_x,"`
	CoordinateAddressY float64    `db:"coordinate_address_y,"`
	CoordinatePointX   float64    `db:"coordinate_point_x,"`
	CoordinatePointY   float64    `db:"coordinate_point_y,"`
	CreatedAt          time.Time  `db:"created_at,"`
	StartedAt          time.Time  `db:"started_at,"`
	DeliveryAt         time.Time  `db:"delivery_at,"`
	CourierId          dto.TypeID `db:"courier_id,"`
	DeliveryStatus     int64      `db:"delivery_status,"`
}
