package service

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"project/helper"
	"project/model"
	"strings"

	"github.com/joho/godotenv"
	"github.com/thanhpk/randstr"
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

	// ENV Configuration
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environtment")
	}

	// Generate Verification Code
	code := randstr.String(20)
	verificationCode := helper.Encode(code)

	//save query
	sqlStatement = `INSERT INTO customer (name, email, password, phone_number, address, verification_code, verified)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = cs.DB.Exec(sqlStatement, customer.Name, customer.Email, customer.Password, customer.PhoneNumber, customer.Address, verificationCode, false)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	firstName := customer.Name
	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[1]
	}

	// Send Email Verification
	emailData := model.EmailData{
		URL:       os.Getenv("CLIENT_ORIGIN") + "/customer/verifyemail/" + code,
		FirstName: firstName,
	}
	sendEmailStatus, err := helper.SendEmail(customer, &emailData)
	if err != nil {
		return sendEmailStatus, err
	}
	return sendEmailStatus, nil
}

func (cs *CustomerServiceImpl) LoginCustomer(customerLogin *model.CustomerLogin) (string, error) {
	
	//check email
	customer, result := cs.IsExistCustomer(customerLogin.Email)
	if !result {
		return "", errors.New("invalid username or password")
	}
	if !customer.Verified {
		return "", errors.New("please verify your email")
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
	sqlStatement := `UPDATE customer SET name = $2, email = $3,  password = $4, phone_number = $5, address = $6
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

func (cs *CustomerServiceImpl) VerifyCustomer(code string) (string, error) {
	// Find Customer by verificationCode
	verificationCode := helper.Encode(code)
	var customer = model.Customer{}
	sqlStatement := `SELECT * FROM customer WHERE verification_code=($1)`
	_, err = cs.DB.Exec(sqlStatement, verificationCode)
	if err != nil {
		return "Invalid verification code or Customer doesn't exists", err
	}

	// check status verified
	if customer.Verified {
		return "Customer already verified", err
	}

	//set verificationCode to empty and update status verified
	sqlStatement = `UPDATE customer SET verification_code = $2, verified = $3 WHERE verification_code = $1;`
	empty := ""
	_, err = cs.DB.Exec(sqlStatement, verificationCode, empty, true)
	if err != nil {
		fmt.Println("error update verify")
		message := fmt.Sprintf("error: %v occurred while updating customer record with id: %d", err, customer.ID)
		return message, err
	}
	return "Email verified successfully", nil
}
