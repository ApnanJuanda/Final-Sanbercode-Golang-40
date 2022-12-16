package service

import (
	"database/sql"
	"errors"
	"fmt"
	"project/model"
)

type CategoryServiceImpl struct {
	DB *sql.DB
}

func NewCategorySevice(DB *sql.DB) CategoryService {
	return &CategoryServiceImpl{
		DB: DB,
	}
}

func (cs *CategoryServiceImpl) AddCategory(addCategory *model.AddCategory) (*model.Category, error) {
	var newCategory = model.Category{}

	if addCategory.Name == "" {
		return nil, errors.New("invalid request body")
	}

	sqlStatement = `INSERT INTO category (name) VALUES ($1) Returning *`
	err = cs.DB.QueryRow(sqlStatement, addCategory.Name).Scan(&newCategory.ID, &newCategory.Name, &newCategory.CreatedAt, &newCategory.UpdatedAt)
	if err != nil {
		return nil, err
	}	

	return &newCategory, nil
}

func (cs *CategoryServiceImpl) GetCategories() (*[]model.Category, error) {
	var categories = []model.Category{}

	sqlStatement = `SELECT * FROM category`

	rows, err := cs.DB.Query(sqlStatement)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category = model.Category{}

		err = rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)

		if err != nil {
			return nil, err
		}

		categories = append(categories, category)

	}
	return &categories, nil
}

func (cs *CategoryServiceImpl) GetCategoryById(id int) (*model.Category, error) {
	var category = model.Category{}
	sqlStatement := `SELECT * FROM category WHERE id=($1)`
	err = cs.DB.QueryRow(sqlStatement, id).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return nil, err		
	}
	return &category, err
}

func (cs *CategoryServiceImpl) GetCategoryByName(name string) (*model.Category, error) {
	var category = model.Category{}
	sqlStatement := `SELECT * FROM category WHERE name=($1)`
	err = cs.DB.QueryRow(sqlStatement, name).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return nil, err		
	}
	return &category, err
}


func (cs *CategoryServiceImpl) UpdateCategoryById(id int, addCategory *model.AddCategory) (int, error) {
	if addCategory.Name == "" {
		return 0, errors.New("invalid request body")
	}
	sqlStatement = `UPDATE category SET name=$2 WHERE id=$1;`
	
	result, err := cs.DB.Exec(sqlStatement, id, addCategory.Name)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while updating category record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := result.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after update data: %v", err)
		return 0, errors.New(e) 
	}

	if count == 0 {
		e := "could not update the category, please try again after sometime"
		return 0, errors.New(e) 
	}
	return int(count), nil
}

func (cs *CategoryServiceImpl) DeleteCategoryById(id int) (int, error) {
	sqlStatement = `DELETE FROM category WHERE id=$1;`
	res, err := cs.DB.Exec(sqlStatement, id)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while delete category record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := res.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after delete data: %v", err)
		return 0, errors.New(e)		
	}

	if count == 0 {
		e := fmt.Sprintf("could not delete the category, there might be no data for ID %d", id)
		return 0, errors.New(e) 
	}
	return int(count), nil
}