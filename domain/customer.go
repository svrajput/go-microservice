package domain

// Lets create structure for Domain Customer
type Customer struct {
	Id          string
	Name        string
	DateOfBirth string
	City        string
	Zipcode     string
	Status      string
}

// Create interface for Customer Repository.
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
