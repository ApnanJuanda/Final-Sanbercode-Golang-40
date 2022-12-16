package model

import "time"

type Feedback struct {
	ID         int       `json:"id,omitempty"`
	CustomerId int       `json:"customer_id,omitempty"`
	ProductId  int       `json:"product_id,omitempty"`
	Review     string    `json:"review,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

type AddFeedback struct {
	ProductId  int       `json:"product_id,omitempty"`
	Review     string    `json:"review,omitempty"`
}
