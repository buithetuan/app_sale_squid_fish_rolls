package models

import "time"

type Orders struct {
	OrderID   uint      `json:"order_id" gorm:"primaryKey;AUTO_INCREMENT"`
	UserID    int       `json:"user_id" gorm:"foreignkey:UserID"`
	CreatedAt time.Time `json:"created_at"`
}
