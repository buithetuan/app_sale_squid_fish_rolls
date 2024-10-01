package models

import "time"

type Carts struct {
	CartID    uint      `json:"cart_id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User  Users       `json:"user" gorm:"foreignKey:UserID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Items []CartItems `json:"items" gorm:"foreignKey:CartID;references:CartID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Khóa ngoại CartID
}
