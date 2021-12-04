package domain

import "github.com/marcelobarreto/banking/errs"

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, *errs.AppError) {
	return s.customers, nil
}

func (s CustomerRepositoryStub) ByID(id string) (*Customer, *errs.AppError) {
	return &Customer{ID: "1", Name: "Marcelo", City: "Toronto", Zipcode: "000", DateOfBirth: "30/06/1994", Status: "active"}, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{ID: "1", Name: "Marcelo", City: "Toronto", Zipcode: "000", DateOfBirth: "1994-06-30", Status: "1"},
		{ID: "2", Name: "Marina", City: "Toronto", Zipcode: "000", DateOfBirth: "1996-07-04", Status: "1"},
	}

	return CustomerRepositoryStub{customers: customers}
}
