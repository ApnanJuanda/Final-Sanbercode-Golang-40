package model

import "time"

type Order struct {
	ID            int       `json:"id,omitempty"`
	CustomerId    int       `json:"customer_id,omitempty"`
	ProductId     int       `json:"product_id,omitempty"`
	Quantity      int       `json:"quantity,omitempty"`
	TotalPrice    int       `json:"total_price,omitempty"`
	TransactionId int       `json:"transactions_id,omitempty"`
	OrderDate     time.Time `json:"order_date,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

type AddOrder struct {
	ProductId int `json:"product_id,omitempty"`
	Quantity  int `json:"quantity,omitempty"`
}

type ResponseOrder struct {
	ID            int       `json:"id,omitempty"`
	Product       Product   `json:"product,omitempty"`
	Quantity      int       `json:"quantity,omitempty"`
	TotalPrice    int       `json:"total_price,omitempty"`
	TransactionId int       `json:"transactions_id,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}
