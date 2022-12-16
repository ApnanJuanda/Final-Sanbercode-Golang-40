package model

import "time"

type Cart struct {
	ID         int       `json:"id,omitempty"`
	CustomerId int       `json:"customer_id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

type AddCart struct {
	CartProducts []AddCartProduct `json:"cart_product,omitempty"`
}

type ResponseCart struct {
	ID           int                   `json:"cart_id,omitempty"`
	CartProducts []ResponseCartProduct `json:"cart_product,omitempty"`
}

type CartProduct struct {
	ID         int       `json:"id,omitempty"`
	ProductId  int       `json:"product_id,omitempty"`
	Quantity   int       `json:"quantity,omitempty"`
	TotalPrice int       `json:"total_price,omitempty"`
	CartId     int       `json:"cart_id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

type AddCartProduct struct {
	ProductId int `json:"product_id,omitempty"`
	Quantity  int `json:"quantity,omitempty"`
}

type ResponseCartProduct struct {
	ID         int       `json:"id,omitempty"`
	Product    Product `json:"product,omitempty"`
	Quantity   int       `json:"quantity,omitempty"`
	TotalPrice int       `json:"total_price,omitempty"`
	CartId     int       `json:"cart_id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
