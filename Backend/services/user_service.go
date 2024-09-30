package services

import (
	"Backend/models"
	"Backend/repo"
	"fmt"
)

type UserService struct {
	userRepo      *repo.UserRepo
	cartRepo      *repo.CartRepo
	orderRepo     *repo.OrderRepo
	cartItemsRepo *repo.CartItemsRepo
}

func NewUserService(userRepo *repo.UserRepo, cartRepo *repo.CartRepo, orderRepo *repo.OrderRepo, cartItemsRepo *repo.CartItemsRepo) *UserService {
	return &UserService{
		userRepo:      userRepo,
		cartRepo:      cartRepo,
		orderRepo:     orderRepo,
		cartItemsRepo: cartItemsRepo,
	}
}

// Authentication
func (s *UserService) UpdateUserService(userID int, email string, phoneNumber string, address string) (*models.Users, error) {
	existingUse, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	if existingUse == nil {
		return nil, fmt.Errorf("user %d does not exists", userID)
	}

	existingUse.Email = email
	existingUse.PhoneNumber = phoneNumber
	existingUse.Address = address

	updatedUser, err := s.userRepo.UpdateUser(userID, existingUse)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *UserService) GetUserService(userID int) (*models.Users, error) {
	return s.userRepo.GetUserByID(userID)
}

func validatePaymentMethod(method string) error {
	validMethods := []string{
		string(models.UnPaid),
		string(models.Tranfer),
		string(models.Cash),
	}

	for _, valid := range validMethods {
		if method == valid {
			return nil
		}
	}
	return fmt.Errorf("invalid shipping status: %s", method)
}
func (s *UserService) UserChoseMethodPayService(userID int, orderID int, payMethod string) error {
	order, err := s.orderRepo.GetOrderByID(orderID)
	if err != nil {
		return fmt.Errorf("failed to get order: %v", err)
	}
	if order == nil {
		return fmt.Errorf("order with ID %d does not exist", orderID)
	}

	if order.UserID != userID {
		return fmt.Errorf("user %d does not have permission to modify order %d", userID, orderID)
	}

	if err := validatePaymentMethod(payMethod); err != nil {
		return err
	}

	err = s.orderRepo.ChosePaymentMethod(order)
}

// Buy product
