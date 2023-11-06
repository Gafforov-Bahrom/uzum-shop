package dto

type Basket struct {
	Id        TypeID
	UserId    TypeID
	ProductId TypeID
	Count     uint64
}

type ListBasketIn struct {
	Query string
}

type ListBasketOut struct {
	Items []*Basket
}

type CreateBasket struct {
	UserId    TypeID
	ProductId TypeID
	Count     uint64
}
