package models

import "time"

type Users struct {
	UserID      uint      `json:"user_id" gorm:"primaryKey;unique"`
	Username    string    `json:"username" gorm:"not null;unique"`
	Email       string    `json:"email"`
	Password    string    `json:"password" gorm:"not null"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Rank        UserRank  `json:"rank" gorm:"default:bronze"`
	TotalBuy    float64   `json:"total_buy" gorm:"default:0.00"`
	IsAdmin     bool      `json:"is_admin" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
