package models

import "time"

// Invoice representa una factura en la base de datos.
type Invoice struct {
	ID          string      `json:"id,omitempty" bson:"_id,omitempty"`
	Number  	string  	`json:"number"`
	Customer    string      `json:"customer,omitempty" bson:"customer,omitempty"`
	Items       []Item      `json:"items,omitempty" bson:"items,omitempty"`
	TotalAmount float64     `json:"totalAmount,omitempty" bson:"totalAmount,omitempty"`
	CreatedAt   time.Time   `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt   time.Time   `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

// Item representa un artículo en una factura.
type Item struct {
	Name     string  `json:"name,omitempty" bson:"name,omitempty"`
	Quantity int     `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Price    float64 `json:"price,omitempty" bson:"price,omitempty"`
}