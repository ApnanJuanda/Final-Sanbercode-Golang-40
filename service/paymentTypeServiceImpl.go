package service

import (
	"database/sql"
	"errors"
	"fmt"
	"project/model"
)

type PaymenTypeServiceImpl struct {
	DB *sql.DB
}

func NewPaymentTypeService(DB *sql.DB) PaymenTypeService {
	return &PaymenTypeServiceImpl{
		DB: DB,
	}
}

func (ps *PaymenTypeServiceImpl) AddPaymenType(addPaymenType *model.AddPaymentType) (*model.PaymentType, error) {
	var newPaymenType = model.PaymentType{}

	if addPaymenType.Name == "" {
		return nil, errors.New("invalid request body")
	}

	sqlStatement = `INSERT INTO payment_type (name) VALUES ($1) Returning *`
	err = ps.DB.QueryRow(sqlStatement, addPaymenType.Name).Scan(&newPaymenType.ID, &newPaymenType.Name, &newPaymenType.CreatedAt, &newPaymenType.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &newPaymenType, nil
}

func (ps *PaymenTypeServiceImpl) GetPaymentTypes() (*[]model.PaymentType, error) {
	var paymentTypes = []model.PaymentType{}

	sqlStatement = `SELECT * FROM payment_type`

	rows, err := ps.DB.Query(sqlStatement)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var paymentType = model.PaymentType{}

		err = rows.Scan(&paymentType.ID, &paymentType.Name, &paymentType.CreatedAt, &paymentType.UpdatedAt)

		if err != nil {
			return nil, err
		}

		paymentTypes = append(paymentTypes, paymentType)

	}
	return &paymentTypes, nil
}

func (ps *PaymenTypeServiceImpl) GetPaymentTypeById(id int) (*model.PaymentType, error) {
	var paymentType = model.PaymentType{}
	sqlStatement := `SELECT * FROM payment_type WHERE id=($1)`
	err = ps.DB.QueryRow(sqlStatement, id).Scan(&paymentType.ID, &paymentType.Name, &paymentType.CreatedAt, &paymentType.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &paymentType, err
}

func (ps *PaymenTypeServiceImpl) GetPaymentTypeByName(name string) (*model.PaymentType, error) {
	var paymentType = model.PaymentType{}
	sqlStatement := `SELECT * FROM payment_type WHERE name=($1)`
	err = ps.DB.QueryRow(sqlStatement, name).Scan(&paymentType.ID, &paymentType.Name, &paymentType.CreatedAt, &paymentType.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &paymentType, err
}

func (ps *PaymenTypeServiceImpl) UpdatePaymentTypeById(id int, addPaymenType *model.AddPaymentType) (int, error) {
	if addPaymenType.Name == "" {
		return 0, errors.New("invalid request body")
	}
	sqlStatement = `UPDATE payment_type SET name=$2 WHERE id=$1;`

	result, err := ps.DB.Exec(sqlStatement, id, addPaymenType.Name)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while updating payment_type record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := result.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after update data: %v", err)
		return 0, errors.New(e)
	}

	if count == 0 {
		e := "could not update the payment_type, please try again after sometime"
		return 0, errors.New(e)
	}
	return int(count), nil
}

func (ps *PaymenTypeServiceImpl) DeletePaymentTypeById(id int) (int, error) {
	sqlStatement = `DELETE FROM payment_type WHERE id=$1;`
	res, err := ps.DB.Exec(sqlStatement, id)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while delete payment_type record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := res.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after delete data: %v", err)
		return 0, errors.New(e)
	}

	if count == 0 {
		e := fmt.Sprintf("could not delete the payment_type, there might be no data for ID %d", id)
		return 0, errors.New(e)
	}
	return int(count), nil
}
