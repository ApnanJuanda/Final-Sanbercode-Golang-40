package model

import "time"

type Product struct {
	ID           int       `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Price        int       `json:"price,omitempty"`
	StockProduct int       `json:"stock_product,omitempty"`
	ImageUrl     string    `json:"image_url,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	CategoryId   int       `json:"category_id,omitempty"`
	SupplierId   int       `json:"supplier_id,omitempty"`
}

type AddProduct struct {
	Name         string    `json:"name,omitempty"`
	Price        int       `json:"price,omitempty"`
	StockProduct int       `json:"stock_product,omitempty"`
	ImageUrl     string    `json:"image_url,omitempty"`
	CategoryId   int       `json:"category_id,omitempty"`
	SupplierId   int       `json:"supplier_id,omitempty"`
}
