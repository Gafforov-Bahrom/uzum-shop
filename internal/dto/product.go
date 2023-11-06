package dto

type TypeID uint64

type Product struct {
	Id          TypeID
	Name        string
	Description string
	Price       uint64
	Count       uint64
}

type ListProductIn struct {
	Limit  uint64
	Offset uint64
	Query  string
}

type ListProductOut struct {
	Count uint64
	Items []*Product
}
