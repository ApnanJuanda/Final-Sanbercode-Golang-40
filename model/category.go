package model

import "time"

type Category struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type AddCategory struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}