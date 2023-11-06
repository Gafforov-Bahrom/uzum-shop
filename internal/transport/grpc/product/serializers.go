package product

import (
	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"
)

func dtoToDesc(in *dto.Product) *desc.Product {
	return &desc.Product{
		Id:          uint64(in.Id),
		Name:        in.Name,
		Description: in.Description,
		Price:       in.Price,
		Count:       in.Count,
	}
}

func listToDesc(in []*dto.Product) []*desc.Product {
	out := make([]*desc.Product, len(in))
	for i, item := range in {
		out[i] = dtoToDesc(item)
	}
	return out
}

func dtoToDescBasket(in *dto.Basket) *desc.Basket {
	return &desc.Basket{
		Id:        uint64(in.Id),
		UserId:    uint64(in.UserId),
		ProductId: uint64(in.ProductId),
		Count:     in.Count,
	}
}

func listToDescBasket(in []*dto.Basket) []*desc.Basket {
	out := make([]*desc.Basket, len(in))
	for i, item := range in {
		out[i] = dtoToDescBasket(item)
	}
	return out
}
