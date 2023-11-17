package shop

import "github.com/Gafforov-Bahrom/uzum_shop/internal/dto"

func dtoToProduct(in *dto.Product) *Product {
	return &Product{
		Id:          uint64(in.Id),
		Name:        in.Name,
		Description: in.Description,
		Count:       in.Count,
		Price:       in.Price,
	}
}

func listToProduct(in []*dto.Product) []*Product {
	out := make([]*Product, len(in))
	for i, item := range in {
		out[i] = dtoToProduct(item)
	}
	return out
}
