package service

import (
	"database/sql"
	"errors"
	"fmt"
	"project/model"
)

type SupplierServiceImpl struct {
	DB *sql.DB
}

func NewSupplierService(DB *sql.DB) SupplierService {
	return &SupplierServiceImpl{
		DB: DB,
	}
}

var (
	err error
	sqlStatement string
)

func (ss *SupplierServiceImpl) AddSupplier(addSupplier *model.AddSupplier) (*model.Supplier, error) {
	var newSupplier = model.Supplier{}

	if addSupplier.Name == "" {
		return nil, errors.New("invalid request body")
	}

	sqlStatement = `INSERT INTO supplier (name) VALUES ($1) Returning *`
	err = ss.DB.QueryRow(sqlStatement, addSupplier.Name).Scan(&newSupplier.ID, &newSupplier.Name, &newSupplier.CreatedAt, &newSupplier.UpdatedAt)
	if err != nil {
		return nil, err
	}	

	return &newSupplier, nil
}

func (ss *SupplierServiceImpl) GetSuppliers() (*[]model.Supplier, error) {
	var suppliers = []model.Supplier{}

	sqlStatement = `SELECT * FROM supplier`

	rows, err := ss.DB.Query(sqlStatement)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var supplier = model.Supplier{}

		err = rows.Scan(&supplier.ID, &supplier.Name, &supplier.CreatedAt, &supplier.UpdatedAt)

		if err != nil {
			return nil, err
		}

		suppliers = append(suppliers, supplier)

	}
	return &suppliers, nil
}

func (ss *SupplierServiceImpl) GetSupplierById(id int) (*model.Supplier, error) {
	var supplier = model.Supplier{}
	sqlStatement := `SELECT * FROM supplier WHERE id=($1)`
	err = ss.DB.QueryRow(sqlStatement, id).Scan(&supplier.ID, &supplier.Name, &supplier.CreatedAt, &supplier.UpdatedAt)
	if err != nil {
		return nil, err		
	}
	return &supplier, err
}

func (ss *SupplierServiceImpl) GetSupplierByName(name string) (*model.Supplier, error) {
	var supplier = model.Supplier{}
	sqlStatement := `SELECT * FROM supplier WHERE name=($1)`
	err = ss.DB.QueryRow(sqlStatement, name).Scan(&supplier.ID, &supplier.Name, &supplier.CreatedAt, &supplier.UpdatedAt)
	if err != nil {
		return nil, err		
	}
	return &supplier, err
}

func (ss *SupplierServiceImpl) UpdateSupplierById(id int, addSupplier *model.AddSupplier) (int, error) {
	if addSupplier.Name == "" {
		return 0, errors.New("invalid request body")
	}
	sqlStatement = `UPDATE supplier SET name=$2 WHERE id=$1;`
	
	result, err := ss.DB.Exec(sqlStatement, id, addSupplier.Name)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while updating supplier record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := result.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after update data: %v", err)
		return 0, errors.New(e) 
	}

	if count == 0 {
		e := "could not update the supplier, please try again after sometime"
		return 0, errors.New(e) 
	}
	return int(count), nil
}

func (ss *SupplierServiceImpl) DeleteSupplierById(id int) (int, error) {
	sqlStatement = `DELETE FROM supplier WHERE id=$1;`
	res, err := ss.DB.Exec(sqlStatement, id)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while delete supplier record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := res.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after delete data: %v", err)
		return 0, errors.New(e)		
	}

	if count == 0 {
		e := fmt.Sprintf("could not delete the supplier, there might be no data for ID %d", id)
		return 0, errors.New(e) 
	}
	return int(count), nil	
}