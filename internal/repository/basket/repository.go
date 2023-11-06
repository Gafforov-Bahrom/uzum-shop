package basket

import (
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

type dbBasket struct {
	Id        uint64 `db:"id"`
	UserId    uint64 `db:"user_id"`
	ProductId uint64 `db:"product_id"`
	Count     uint64 `db:"count"`
}

func (item dbBasket) toDTO() *dto.Basket {
	return &dto.Basket{
		Id:        dto.TypeID(item.Id),
		UserId:    dto.TypeID(item.UserId),
		ProductId: dto.TypeID(item.ProductId),
		Count:     item.Count,
	}
}

func ListToDto(items []dbBasket) []*dto.Basket {
	response := make([]*dto.Basket, len(items))
	for i, item := range items {
		response[i] = item.toDTO()
	}
	return response
}
