package repo

import (
	"Backend/models"
	"fmt"
	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) GetOrderByID(orderID int) (*models.Orders, error) {
	var order models.Orders
	if err := r.db.Where("order_id = ?", orderID).First(&order).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("order not found")
		}
		return nil, fmt.Errorf("error fetching order: %v", err)
	}
	return &order, nil
}

func (r *OrderRepo) UpdateStatusOrder(order *models.Orders) error {
	if err := r.db.Save(order).Error; err != nil {
		return fmt.Errorf("error updating status order: %v", err)
	}
	return nil
}

func (r *OrderRepo) ChosePaymentMethod(order *models.Orders) error {
	if err := r.db.Save(order).Error; err != nil {
		return fmt.Errorf("error chosing payment method order: %v", err)
	}
	return nil
}

func (r *OrderRepo) CreateOrder(cartID int)
