package service

import "project/model"

type CategoryService interface {
	AddCategory(addCategory *model.AddCategory) (*model.Category, error)
	GetCategories() (*[]model.Category, error)
	GetCategoryById(id int) (*model.Category, error)
	GetCategoryByName(name string) (*model.Category, error)
	UpdateCategoryById(id int, addCategory *model.AddCategory) (int, error)
	DeleteCategoryById(id int) (int, error)
}