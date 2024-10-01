package repo

import (
	"Backend/models"
	"fmt"
	"gorm.io/gorm"
)

type CartItemsRepo struct {
	db *gorm.DB
}

func NewCartItemRepo(db *gorm.DB) *CartItemsRepo {
	return &CartItemsRepo{db: db}
}

func (r *CartItemsRepo) getCartItem(cartID uint, productID string) (*models.CartItems, error) {
	var cartItem models.CartItems
	result := r.db.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&cartItem)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &cartItem, nil
}

func (r *CartItemsRepo) updateCartItemQuantity(cartItem *models.CartItems, quantity int) error {
	cartItem.Quantity += quantity
	if err := r.db.Save(cartItem).Error; err != nil {
		return fmt.Errorf("failed to update cart item quantity: %v", err)
	}
	return nil
}

func (r *CartItemsRepo) addNewCartItem(cartID uint, productID string, quantity int) error {
	cartItem := models.CartItems{
		CartID:    cartID,
		ProductID: productID,
		Quantity:  quantity,
	}
	if err := r.db.Create(&cartItem).Error; err != nil {
		return fmt.Errorf("failed to add new item to cart: %v", err)
	}
	return nil
}

func (r *CartItemsRepo) AddItemToCart(cartID uint, productID string, quantity int) error {
	existingCartItem, err := r.getCartItem(cartID, productID)
	if err != nil {
		return fmt.Errorf("failed to check cart item: %v", err)
	}

	if existingCartItem != nil {
		if err := r.updateCartItemQuantity(existingCartItem, quantity); err != nil {
			return err
		}
	} else {
		if err := r.addNewCartItem(cartID, productID, quantity); err != nil {
			return err
		}
	}

	return nil
}

func (r *CartItemsRepo) RemoveItemFromCart(cartID int, productID string) error {
	err := r.db.Where("cart_id = ? AND product_id = ?", cartID, productID).Delete(&models.CartItems{}).Error
	return err
}

func (r *CartItemsRepo) ClearCartItems(cartID int) error {
	if err := r.db.Where("cart_id = ?", cartID).Delete(&models.CartItems{}).Error; err != nil {
		return fmt.Errorf("failed to clear cart items: %v", err)
	}
	return nil
}
