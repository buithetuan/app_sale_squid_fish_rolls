package repo

import (
	"Backend/models"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type CartRepo struct {
	db *gorm.DB
}

func NewCartRepo(db *gorm.DB) *CartRepo {
	return &CartRepo{db: db}
}

func (r *CartRepo) CreateCart(cart *models.Carts) error {
	if err := r.db.Create(cart).Error; err != nil {
		return fmt.Errorf("CartRepo CreateCart err: %v", err)
	}
	return nil
}

func (r *CartRepo) GetCartByUserID(ctx context.Context, userID int) (*models.Carts, error) {
	var cart models.Carts

	err := r.db.WithContext(ctx).Preload("CartItems.Products").Where("user_id = ?", userID).Find(&cart).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &cart, nil
}
