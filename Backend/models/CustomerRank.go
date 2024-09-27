package models

import "time"

type CustomerRank struct {
	RankID             uint      `json:"rank_id" gorm:"primaryKey"`
	RankName           string    `json:"rank_name"`
	DiscountPercentage float64   `json:"discount_percentage"`
	MinTotalPurchase   float64   `json:"min_total_purchase"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
