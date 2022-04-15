package repository

type Address struct{}

type AddressRespository interface {
	GetAll()
	GetById()
	GetByCustomerID()
	Create()
	UpdateByID()
	DeleteByID()
}
