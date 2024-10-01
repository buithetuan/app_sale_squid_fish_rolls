package models

import "time"

type OrderItems struct {
	OrderItemsID uint      `json:"order_items_id" gorm:"primaryKey"`
	OrderID      uint      `json:"order_id"`
	ProductID    string    `json:"product_id"`
	ProductName  string    `json:"product_name"`
	Quantity     int       `json:"quantity"`
	Price        float64   `json:"price"`
	TotalPrice   float64   `json:"total_price"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
