package dto

type User struct {
	Id                 TypeID
	Name               string
	Surname            string
	Phone              string
	Login              string
	PasswordHash       string
	Role               string
	Address            string
	CoordinateAddressX float64
	CoordinateAddressY float64
}
