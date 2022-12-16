package model

import "time"

type Payment struct {
	ID            int       `json:"id,omitempty"`
	PaymentTypeId string    `json:"payment_type_id,omitempty"`
	PaymenStatus  string    `json:"payment_status,omitempty"`
	TotalPayment  int       `json:"total_payment,omitempty"`
	TransactionId int       `json:"transactions_id,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

type AddPayment struct {
	PaymentTypeId string `json:"payment_type_id,omitempty"`
}
