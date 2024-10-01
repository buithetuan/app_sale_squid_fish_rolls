package repo

import (
	"Backend/models"
	"gorm.io/gorm"
)

type OrderItemsRepo struct {
	db *gorm.DB
}

func NewOrderItems(db *gorm.DB) *OrderItemsRepo {
	return &OrderItemsRepo{db: db}
}
func (r *OrderItemsRepo) CreateOrderItemsFromCart(cartItems []models.CartItems, orderID uint) (float64, error) {
	totalAmount := 0.0
	for _, item := range cartItems {
		orderItem := models.OrderItems{
			OrderID:     orderID,
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			Quantity:    item.Quantity,
			Price:       item.Price,
			TotalPrice:  float64(item.Quantity) * item.Price,
		}

		if err := r.db.Create(&orderItem).Error; err != nil {
			return 0.0, err
		}
		totalAmount += orderItem.TotalPrice
	}
	return totalAmount, nil
}
