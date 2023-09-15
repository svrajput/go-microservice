package domain

// TODO: create structure CustomerRepositoryStub has variable customers
type CustomerRepositoryStub struct {
	customers []Customer
}

// TODO: Implement FindAll() of receiver type CustomerRepositoryStub
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// TODO: Heler function to return CustomerRepositoryStub object. - dummy customer record.
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "John", "1980-01-10", "Dallas", "111111", "1"},
		{"2001", "Jason", "1985-02-15", "Dallas", "111111", "1"},
		{"3001", "Matt", "1990-02-20", "Dallas", "111111", "1"},
	}

	return CustomerRepositoryStub{customers}
}
