package model

import "time"

type Transaction struct {
	ID         int       `json:"id,omitempty"`
	CustomerId string    `json:"customer_id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

type AddTransaction struct {
	Orders         []AddOrder `json:"order,omitempty"`
	PaymentTypeId  int        `json:"payment_type,omitempty"`
	CourierCompany string     `json:"courier_company,omitempty"`
}

type ResponseTransaction struct {
	ID       int             `json:"tx_id,omitempty"`
	Orders   []ResponseOrder `json:"order,omitempty"`
	Payment  Payment         `json:"payment,omitempty"`
	Delivery Delivery        `json:"delivery,omitempty"`
}
