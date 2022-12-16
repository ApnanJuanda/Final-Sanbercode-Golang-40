package service

import "project/model"

type SupplierService interface {
	AddSupplier(addSupplier *model.AddSupplier) (*model.Supplier, error)
	GetSuppliers() (*[]model.Supplier, error)
	GetSupplierById(id int) (*model.Supplier, error)
	GetSupplierByName(name string) (*model.Supplier, error)
	UpdateSupplierById(id int, addSupplier *model.AddSupplier) (int, error)
	DeleteSupplierById(id int) (int, error)
}
