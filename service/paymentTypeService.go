package service

import "project/model"

type PaymenTypeService interface {
	AddPaymenType(addPaymenType *model.AddPaymentType) (*model.PaymentType, error)
	GetPaymentTypes() (*[]model.PaymentType, error)
	GetPaymentTypeById(id int) (*model.PaymentType, error)
	GetPaymentTypeByName(name string) (*model.PaymentType, error)
	UpdatePaymentTypeById(id int, addPaymenType *model.AddPaymentType) (int, error)
	DeletePaymentTypeById(id int) (int, error)
}