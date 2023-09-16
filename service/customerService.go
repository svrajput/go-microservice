package service

import "github.com/svrajput/go-microservice/domain"

// TODO - Create CustomerService Interface
// defined GetAllCustomers - which returns list of customers.
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

// TODO - create DefaultCustomerService Struct has dependency of domain.customerRepository interface
type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

// Implement GetAllCustomers
func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repository.FindAll()
}

// TODO create Helper function NewCustomerService
// To create instance of DefaultCustomer service and inject dependency
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
