package model

import "time"

type Supplier struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type AddSupplier struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
