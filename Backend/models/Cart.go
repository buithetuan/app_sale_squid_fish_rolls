package models

import "time"

type Cart struct {
	CartID    uint      `json:"cart_id" gorm:"primaryKey"`
	UserID    int       `json:"user_id"`
	ProductId string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
