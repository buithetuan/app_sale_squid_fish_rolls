package services

import (
	"Backend/dto"
	"Backend/models"
	"Backend/repo"
	"context"
	"fmt"
)

type UserService struct {
	userRepo       *repo.UserRepo
	cartRepo       *repo.CartRepo
	orderRepo      *repo.OrderRepo
	cartItemsRepo  *repo.CartItemsRepo
	orderItemsRepo *repo.OrderItemsRepo
}

func NewUserService(userRepo *repo.UserRepo, cartRepo *repo.CartRepo, orderRepo *repo.OrderRepo, cartItemsRepo *repo.CartItemsRepo, orderItemsRepo *repo.OrderItemsRepo) *UserService {
	return &UserService{
		userRepo:       userRepo,
		cartRepo:       cartRepo,
		orderRepo:      orderRepo,
		cartItemsRepo:  cartItemsRepo,
		orderItemsRepo: orderItemsRepo,
	}
}

// Authentication
func (s *UserService) UpdateUserService(userID uint, email string, phoneNumber string, address string) (*models.Users, error) {
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

func (s *UserService) GetUserService(userID uint) (*models.Users, error) {
	return s.userRepo.GetUserByID(userID)
}

// User Control Item Cart
func (s *UserService) UserAddItemsToCart(ctx context.Context, userID uint, productID string, quantity int) error {
	cart, err := s.cartRepo.GetCartByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("cannot get cart for user %d: %v", userID, err)
	}
	if cart == nil {
		return fmt.Errorf("no cart found for user ID %d", userID)
	}
	err = s.cartItemsRepo.AddItemToCart(cart.CartID, productID, quantity)
	if err != nil {
		return fmt.Errorf("failed to add item to cart: %v", err)
	}

	return nil
}

func (s *UserService) UserRemoveItemFromCart(ctx context.Context, userID uint, productID string) error {
	cart, err := s.cartRepo.GetCartByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("cannot get cart for user %d: %v", userID, err)
	}
	if cart == nil {
		return fmt.Errorf("no cart found for user ID %d", userID)
	}
	err = s.cartItemsRepo.RemoveItemFromCart(int(cart.CartID), productID)
	if err != nil {
		return fmt.Errorf("failed to remove item from cart: %v", err)
	}
	return nil
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

func (s *UserService) UserOrderService(ctx context.Context, userID uint, paymentMethod string, deliveryAddress string) (*dto.OrderResponse, error) {
	if err := validatePaymentMethod(string(paymentMethod)); err != nil {
		return nil, err
	}
	if deliveryAddress == "" {
		address, err := s.orderRepo.GetDefaultAddress(userID)
		if err != nil {
			return nil, fmt.Errorf("error retrieving default address: %v", err)
		}
		deliveryAddress = address
	}
	cart, err := s.cartRepo.GetCartByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("cannot get cart for user %d: %v", userID, err)
	}

	if cart == nil {
		return nil, fmt.Errorf("no cart found for user ID %d", userID)
	}

	rank, err := s.userRepo.GetRankFromUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user rank: %v", err)
	}

	totalAmount := 0.0
	for _, item := range cart.Items {
		totalAmount += item.Price
	}
	discountRate := models.GetDiscountRate(rank)
	discountAmount := totalAmount * discountRate
	finalAmount := totalAmount - discountAmount

	order := models.Orders{
		UserID:          userID,
		PaymentMethod:   models.PaymentMethod(paymentMethod),
		DeliveryAddress: deliveryAddress,
		TotalPrice:      totalAmount,
		DiscountApplied: discountAmount,
		FinalPrice:      finalAmount,
		OrderStatus:     models.Pending,
	}

	orderID, err := s.orderRepo.CreateOrder(&order)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %v", err)
	}

	orderItems := []dto.CartItemsResponse{}
	for _, item := range cart.Items {
		orderItems = append(orderItems, dto.CartItemsResponse{
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			Quantity:    item.Quantity,
			Price:       fmt.Sprintf("%.2f", item.Price),
			TotalPrice:  fmt.Sprintf("%.2f", item.Price*float64(item.Quantity)),
		})
	}

	if _, err := s.orderItemsRepo.CreateOrderItemsFromCart(cart.Items, orderID); err != nil {
		return nil, fmt.Errorf("failed to create order items: %v", err)
	}

	if err := s.cartItemsRepo.ClearCartItems(int(cart.CartID)); err != nil {
		return nil, fmt.Errorf("failed to clear cart items: %v", err)
	}

	orderResponse := &dto.OrderResponse{
		OrderID:         orderID,
		UserID:          userID,
		OrderItems:      orderItems,
		TotalPrice:      finalAmount,
		ShipPrice:       0,
		DiscountPrice:   discountAmount,
		DeliveryAddress: deliveryAddress,
		PaymentMethod:   string(paymentMethod),
	}

	return orderResponse, nil
}
