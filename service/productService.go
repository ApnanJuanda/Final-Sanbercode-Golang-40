package service

import "project/model"

type ProductService interface {
	AddProduct(addProduct *model.AddProduct) (*model.Product, error)
	GetProducts() (*[]model.Product, error)
	GetProductById(id int) (*model.Product, error)
	UpdateProductById(id int, addProduct *model.AddProduct) (int, error)
	DeleteProductById(id int) (int, error)
}