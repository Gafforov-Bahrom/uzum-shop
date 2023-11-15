package shop

type Product struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Count       uint64 `json:"count"`
	Price       uint64 `json:"price"`
}

type AddProductRequest struct {
	ProductId uint64 `json:"product_id"`
	UserId    uint64 `json:"user_id"` // Temp solution, need to get user_id from authorization token
	Count     uint64 `json:"count"`
}

type AddProductResponse struct {
}
