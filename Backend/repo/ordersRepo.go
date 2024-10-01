package repo

import (
	"Backend/models"
	"fmt"
	"gorm.io/gorm"
)

type OrderRepo struct {
	db       *gorm.DB
	userRepo *UserRepo
}

func NewOrderRepo(db *gorm.DB, userRepo *UserRepo) *OrderRepo {
	return &OrderRepo{db: db, userRepo: userRepo}
}

func (r *OrderRepo) GetOrderByID(orderID uint) (*models.Orders, error) {
	var order models.Orders
	if err := r.db.Preload("Items").First(&order, orderID).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepo) UpdateStatusOrder(order *models.Orders) error {
	if err := r.db.Save(order).Error; err != nil {
		return fmt.Errorf("error updating status order: %v", err)
	}
	return nil
}

func (r *OrderRepo) GetDefaultAddress(userID uint) (string, error) {
	user, err := r.userRepo.GetUserByID(userID)
	if err != nil {
		return "", err
	}
	return user.Address, nil
}

func (r *OrderRepo) CreateOrder(order *models.Orders) (uint, error) {
	if err := r.db.Create(order).Error; err != nil {
		return 0, err
	}
	return order.OrderID, nil
}
