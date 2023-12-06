package shop

type Product struct {
	Id          uint64 `json:"-"`
	Name        string `json:"name"`
	Description string `json:"-"`
	Count       uint64 `json:"-"`
	Price       uint64 `json:"price"`
}

type GetProductsResponse struct {
	Count    uint64     `json:"count"`
	Products []*Product `json:"products"`
}

type AddProductRequest struct {
	ProductId uint64 `json:"product_id"`
	Count     uint64 `json:"count"`
}

type UpdateProductRequest struct {
	ProductId uint64 `json:"product_id"`
	Count     uint64 `json:"count"`
}

type CancelOrderRequest struct {
	OrderId uint64 `json:"user_id"`
}
