package model

import "time"

type Delivery struct {
	ID             int       `json:"id,omitempty"`
	CourierCompany string    `json:"courier_company,omitempty"`
	StatusDelivery string    `json:"status_delivery,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	TransactionId  int       `json:"transactions_id,omitempty"`
}
