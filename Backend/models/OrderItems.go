package models

import "time"

type OrderItems struct {
	OrderItemsID uint      `json:"order_items_id" gorm:"primaryKey"`
	OrderID      int       `json:"order_id"`
	ProductID    int       `json:"product_id"`
	ProductName  string    `json:"product_name"`
	Quantity     int       `json:"quantity"`
	Price        float64   `json:"price"`
	TotalPrice   float64   `json:"total_price"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	//CartItems CartItems `json:"cart_items" gorm:"foreignKey:OrderItemID"`
}
