package service

import (
	"database/sql"
	"errors"
	"fmt"
	"project/helper"
	"project/model"
)

type CustomerServiceImpl struct {
	DB *sql.DB
}

func NewCustomerService(DB *sql.DB) CustomerService {
	return &CustomerServiceImpl{
		DB: DB,
	}
}

func (cs *CustomerServiceImpl) RegisterCustomer(customerRegister *model.CustomerRegister) (string, error) {
	//populator
	customer, err := helper.CustomerPopulator(customerRegister)
	if err != nil {
		return "", err
	}

	//check email sudah ada atau belum
	_, result := cs.IsExistCustomer(customerRegister.Email)
	if result {
		return "", errors.New("email has been registered")
	}

	//save query
	sqlStatement = `INSERT INTO customer (name, email, password, phone_number, address)
	VALUES ($1, $2, $3, $4, $5)`
	_, err = cs.DB.Exec(sqlStatement, customer.Name, customer.Email, customer.Password, customer.PhoneNumber, customer.Address)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	
	token, err := customer.GenerateJWT()
	if err != nil {
		return "", err
	}
	return token, nil
}

func (cs *CustomerServiceImpl) LoginCustomer(customerLogin *model.CustomerLogin) (string, error) {
	
	//check email
	customer, result := cs.IsExistCustomer(customerLogin.Email)
	if !result {
		return "", errors.New("invalid username or password")
	}

	//check password
	if correctPassword := customer.CorrectPassword(customerLogin.Password); !correctPassword {
		return "", errors.New("invalid username or password")
	}
	
	token, err := customer.GenerateJWT()
	if err != nil {
		return "", errors.New("error when generate token")
	}
	return token, nil
	
}

func (cs *CustomerServiceImpl) IsExistCustomer(email string) (*model.Customer, bool) {
	var customer = model.Customer{}
	sqlStatement := `SELECT * FROM customer WHERE email=($1)`
	err = cs.DB.QueryRow(sqlStatement, email).Scan(&customer.ID, &customer.Email, &customer.Password, &customer.Name,
	&customer.PhoneNumber, &customer.Address, &customer.CreatedAt, &customer.UpdatedAt)
	if err == nil {
		return &customer, true
	} else {
		return &customer, false
	}
}

func (cs *CustomerServiceImpl) UpdateCustomer(id int, customerRegister *model.CustomerRegister) (int, error) {
	//populator
	customer, err := helper.CustomerPopulator(customerRegister)
	if err != nil {
		return 0, err
	}

	//save query
	sqlStatement := `UPDATE product SET name = $2, email = $3,  password = $4, phone_number = $5, address = $6
	WHERE id = $1;`
	result, err := cs.DB.Exec(sqlStatement, id, customer.Name, customer.Email, customer.Password, customer.PhoneNumber, customer.Address)

	if err != nil {
		e := fmt.Sprintf("error: %v occurred while updating customer record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := result.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after update data: %v", err)
		return 0, errors.New(e) 
	}

	if count == 0 {
		e := "could not update the customer, please try again after sometime"
		return 0, errors.New(e) 
	}
	return int(count), nil	
}

func (cs *CustomerServiceImpl) DeleteCustomer(id int) (int, error) {
	sqlStatement = `DELETE FROM customer WHERE id=$1;`
	res, err := cs.DB.Exec(sqlStatement, id)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while delete customer record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := res.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after delete data: %v", err)
		return 0, errors.New(e)		
	}

	if count == 0 {
		e := fmt.Sprintf("could not delete the customer, there might be no data for ID %d", id)
		return 0, errors.New(e) 
	}
	return int(count), nil
}
