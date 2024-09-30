package models

import "time"

type Orders struct {
	OrderID         uint          `json:"order_id" gorm:"primaryKey"`
	UserID          int           `json:"user_id" gorm:"foreignKey:UserID"`
	ShipPrice       int           `json:"ship_price"`
	TotalPrice      float64       `json:"total_price"`
	DiscountApplied float64       `json:"discount_applied"`
	FinalPrice      float64       `json:"final_price"`
	OrderStatus     OrderStatus   `json:"order_status" default:"pending"`
	PaymentMethod   PaymentMethod `json:"payment_method" default:"unpaid"`
	PaymentStatus   PaymentStatus `json:"payment_status" default:"pending"`
	ShipStatus      ShipStatus    `json:"ship_status" default:"pending"`
	DeliveryAddress string        `json:"delivery_address"`
	OrderDate       time.Time     `json:"order_date"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}
