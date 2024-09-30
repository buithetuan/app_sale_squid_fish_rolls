package repo

import (
	"gorm.io/gorm"
)

type CartItemsRepo struct {
	db *gorm.DB
}

func NewCartItemRepo(db *gorm.DB) *CartItemsRepo {
	return &CartItemsRepo{db: db}
}

//func (r *CartItemRepo) AddItemsToCart(productID string, quantity int) (*models.CartItems, error) {
//
//}
//
//func (r *CartItemRepo) RemoveItemsFromCart(productID string) (*models.CartItems, error) {
//
//}
//
//func (r *CartItemRepo) UpdatedItemsFromCart(productID string, quantity int) (*models.CartItems, error) {
//
//}
