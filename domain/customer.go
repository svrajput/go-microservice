package domain

// Lets create structure for Domain Customer
type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

// Create interface for Customer Repository.
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindById(id string) (*Customer, error)
}
