package product

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

type dbProduct struct {
	Id          uint64 `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Price       uint64 `db:"price"`
	Count       uint64 `db:"count"`
}

func (item dbProduct) toDTO() *dto.Product {
	return &dto.Product{
		Id:          dto.TypeID(item.Id),
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		Count:       item.Count,
	}
}

func listToDto(items []dbProduct) []*dto.Product {
	response := make([]*dto.Product, len(items))
	for i, item := range items {
		response[i] = item.toDTO()
	}
	return response
}
