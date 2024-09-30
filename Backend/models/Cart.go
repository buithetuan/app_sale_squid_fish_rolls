package models

import "time"

type Carts struct {
	CartID    uint      `json:"cart_id" gorm:"primaryKey"`
	UserID    int       `json:"user_id" gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User  Users       `json:"users" gorm:"foreignKey:UserID"`
	Items []CartItems `json:"cart_items" gorm:"foreignKey:CartID"`
}
