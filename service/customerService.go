package service

import "project/model"

type CustomerService interface {
	RegisterCustomer(customerRegister *model.CustomerRegister) (string, error)
	VerifyCustomer(code string) (string, error)
	LoginCustomer(customerLogin *model.CustomerLogin) (string, error)
	IsExistCustomer(email string) (*model.Customer, bool)
	UpdateCustomer(id int, customerRegister *model.CustomerRegister) (int, error)
	DeleteCustomer(id int) (int, error)
}