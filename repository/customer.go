package repository

type Customer struct {
	CustomerID int    `db:"customer_id"`
	Name       string `db:"name"`
	DateBirth  string `db:"date_of_birth"`
	City       string `db:"city"`
	ZipCode    string `db:"zipcode"`
	Status     int    `db:"status"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(int) (*Customer, error)
}
