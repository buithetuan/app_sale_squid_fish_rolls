package models

import "time"

type Product struct {
	ProductID string    `json:"product_id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Note      string    `json:"note"`
	Price     float64   `json:"price" gorm:"not null"`
	Stock     int       `json:"stock" gorm:"default:0"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
