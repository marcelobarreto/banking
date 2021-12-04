package service

import (
	"github.com/marcelobarreto/banking/domain"
	"github.com/marcelobarreto/banking/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, *errs.AppError)
	GetCustomer(id string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ByID(id)
}

func NewCustomerService(repository domain.CustomerRepository) CustomerService {
	return DefaultCustomerService{repo: repository}
}
