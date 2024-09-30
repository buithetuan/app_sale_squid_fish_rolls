package models

import "time"

type CartItems struct {
	CartItemID  uint      `json:"cart_item_id" gorm:"primaryKey"`
	CartID      int       `json:"cart_id" gorm:"not null"`
	ProductID   string    `json:"product_id" gorm:"not null"`
	ProductName string    `json:"product_name" gorm:"not null"`
	Price       float64   `json:"price" gorm:"not null"`
	Quantity    int       `json:"quantity" gorm:"default:1"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Products Products `json:"product" gorm:"foreignKey:ProductID"`
	Cart     Carts    `json:"cart" gorm:"foreignKey:CartID"`
}
