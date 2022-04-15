package repository

type Customer struct {
	CustomerID int    `db:"customer_id"`
	FirstName  string `db:"first_name"`
	LastName   string `db:"last_name"`
	Address
}
type CustomerRepository interface {
	GetAll()
	GetById(id int)
	Create()
	UpdateByID()
	DeleteByID()
}
