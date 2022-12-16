package service

import (
	"database/sql"
	"errors"
	"fmt"
	"project/model"
)

type TransactionServiceImpl struct {
	DB *sql.DB
}

func NewTransactionService(DB *sql.DB) TransactionService {
	return &TransactionServiceImpl{
		DB: DB,
	}
}

func (ts *TransactionServiceImpl) AddTransaction(addTransaction *model.AddTransaction, customerId int) (*model.ResponseTransaction, error) {
	var responseTransaction = model.ResponseTransaction{}
	var newTransaction = model.Transaction{}

	if addTransaction.Orders == nil {
		return nil, errors.New("order can not empty")
	}

	// save transaction first to get tx_id
	sqlStatement = `INSERT INTO transactions (customer_id) VALUES ($1) Returning *`
	err = ts.DB.QueryRow(sqlStatement, customerId).Scan(&newTransaction.ID, &newTransaction.CustomerId,
		&newTransaction.CreatedAt, &newTransaction.UpdatedAt)
	if err != nil {
		return nil, err
	}

	txId := newTransaction.ID

	// save masing-masing order
	err = ts.SaveOrder(addTransaction, txId, customerId)
	if err != nil {
		return nil, err
	}

	// save delivery
	delivery, err := ts.SaveDelivery(addTransaction, txId)
	if err != nil {
		return nil, err
	}

	// save payment
	payment, err := ts.SavePayment(addTransaction, txId)
	if err != nil {
		return nil, err
	}

	responseOrders, err := ts.GetOrderByTxid(txId)
	if err != nil {
		return nil, err
	}

	responseTransaction.ID = newTransaction.ID
	responseTransaction.Orders = *responseOrders
	responseTransaction.Delivery = *delivery
	responseTransaction.Payment = *payment

	return &responseTransaction, nil

}

func (ts *TransactionServiceImpl) SaveOrder(addTransaction *model.AddTransaction, txId int, customerId int) error {
	//save masing-masing order dengan txId yang baru disimpan
	orders := addTransaction.Orders
	for _, addOrder := range orders {
		var order = model.Order{}

		// get productById -> dapatin harganya -> kali dengan quantitynya
		productService := NewProductService(ts.DB)
		product, _ := productService.GetProductById(addOrder.ProductId)
		if product == nil {
			return errors.New("product_id is not valid")
		}

		// populate order
		if addOrder.Quantity > product.StockProduct {
			return errors.New("stock Product is not enough")
		}
		order.CustomerId = customerId
		order.ProductId = addOrder.ProductId
		order.Quantity = addOrder.Quantity
		order.TotalPrice = addOrder.Quantity * product.Price
		order.TransactionId = txId

		// save query
		sqlStatement = `INSERT INTO orders (customer_id, product_id, quantity, total_price, transactions_id) VALUES ($1, $2, $3, $4, $5)`
		_, err = ts.DB.Exec(sqlStatement, order.CustomerId, order.ProductId, order.Quantity, order.TotalPrice, order.TransactionId)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}

	return nil
}

func (ts *TransactionServiceImpl) SaveDelivery(addTransaction *model.AddTransaction, txId int) (*model.Delivery,error) {
	var newDelivery = model.Delivery{}
	status := "PREPARED"

	if addTransaction.CourierCompany == "" {
		return nil, errors.New("courier_company is required")
	}

	// save query
	sqlStatement = `INSERT INTO delivery (courier_company, status_delivery, transactions_id) VALUES ($1, $2, $3) Returning *`
	err = ts.DB.QueryRow(sqlStatement, addTransaction.CourierCompany, status, txId).Scan(&newDelivery.ID, &newDelivery.CourierCompany, 
		&newDelivery.StatusDelivery, &newDelivery.CreatedAt, &newDelivery.UpdatedAt, &newDelivery.TransactionId)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return &newDelivery, nil
}

func (ts *TransactionServiceImpl) SavePayment(addTransaction *model.AddTransaction, txId int) (*model.Payment, error) {
	var newPayment = model.Payment{}
	status := "NOT_PAID"


	// validasi payment_type_id
	paymentTypeService := NewPaymentTypeService(ts.DB)
	paymentType, _ := paymentTypeService.GetPaymentTypeById(addTransaction.PaymentTypeId)
	if paymentType == nil {
		return nil, errors.New("payment_type_id is not valid")
	}

	// Count totalPayment
	totalPayment := 0
	responseOrders, err := ts.GetOrderByTxid(txId)
	if err != nil {
		return nil, err
	}

	for _, value := range *responseOrders {
		totalPayment += value.TotalPrice
	}

	// save query
	sqlStatement = `INSERT INTO payment (payment_type_id, payment_status, total_payment, transactions_id) VALUES ($1, $2, $3, $4) Returning *`
	err = ts.DB.QueryRow(sqlStatement, addTransaction.PaymentTypeId, status, totalPayment, txId).Scan(&newPayment.ID, &newPayment.PaymentTypeId, &newPayment.PaymenStatus, 
		&newPayment.TotalPayment, &newPayment.TransactionId, &newPayment.CreatedAt, &newPayment.UpdatedAt)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return &newPayment, nil
}

func (ts *TransactionServiceImpl) GetTransaction(txId int) (*model.ResponseTransaction, error) {
	var responseTransaction = model.ResponseTransaction{}

	responseOrders, err := ts.GetOrderByTxid(txId)
	if err != nil {
		return nil, err
	}	
	
	delivery, err := ts.GetDeliveryByTxid(txId)
	if err != nil {
		return nil, err
	}

	payment, err := ts.GetPaymentByTxid(txId)
	if err != nil {
		return nil, err
	}

	responseTransaction.ID = txId
	responseTransaction.Orders = *responseOrders
	responseTransaction.Delivery = *delivery
	responseTransaction.Payment = *payment

	return &responseTransaction, nil
}

func (ts *TransactionServiceImpl) GetOrderByTxid(txId int) (*[]model.ResponseOrder, error) {
	var orders = []model.Order{}

	sqlStatement = `SELECT * FROM orders WHERE transactions_id=($1)`

	rows, err := ts.DB.Query(sqlStatement, txId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order = model.Order{}

		err = rows.Scan(&order.ID, &order.CustomerId, &order.ProductId, &order.Quantity, &order.TotalPrice,
			&order.TransactionId, &order.OrderDate, &order.UpdatedAt)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)

	}

	// populate
	var responseOrders = []model.ResponseOrder{}
	for _, value := range orders {
		var responseOrder = model.ResponseOrder{}
		responseOrder.ID = value.ID
		responseOrder.Quantity = value.Quantity
		responseOrder.TotalPrice = value.TotalPrice
		responseOrder.TransactionId = value.TransactionId

		//get product by id
		productService := NewProductService(ts.DB)
		product, _ := productService.GetProductById(value.ProductId)
		if err != nil {
			return nil, err
		}
		responseOrder.Product = *product
		responseOrders = append(responseOrders, responseOrder)
	}
	return &responseOrders, nil
}

func (ts *TransactionServiceImpl) GetDeliveryByTxid(txId int) (*model.Delivery, error) {
	var delivery = model.Delivery{}
	sqlStatement := `SELECT * FROM delivery WHERE transactions_id=($1)`
	err = ts.DB.QueryRow(sqlStatement, txId).Scan(&delivery.ID, &delivery.CourierCompany, 
		&delivery.StatusDelivery, &delivery.CreatedAt, &delivery.UpdatedAt, &delivery.TransactionId)
	if err != nil {
		return nil, err		
	}
	return &delivery, nil
}

func (ts *TransactionServiceImpl) GetPaymentByTxid(txId int) (*model.Payment, error) {
	var payment = model.Payment{}
	sqlStatement := `SELECT * FROM payment WHERE transactions_id=($1)`
	err = ts.DB.QueryRow(sqlStatement, txId).Scan(&payment.ID, &payment.PaymentTypeId, &payment.PaymenStatus, 
		&payment.TotalPayment, &payment.TransactionId, &payment.CreatedAt, &payment.UpdatedAt)
	if err != nil {
		return nil, err		
	}
	return &payment, nil
}