package services

import (
	"Backend/models"
	"Backend/repo"
	"fmt"
)

type AdminService struct {
	userRepo    *repo.UserRepo
	productRepo *repo.ProductRepo
	orderRepo   *repo.OrderRepo
}

func NewAdminService(userRepo *repo.UserRepo, productRepo *repo.ProductRepo, orderRepo *repo.OrderRepo) *AdminService {
	return &AdminService{
		userRepo:    userRepo,
		productRepo: productRepo,
		orderRepo:   orderRepo,
	}
}

func (as *AdminService) CheckAdmin(userID int) error {
	isAdmin, err := as.userRepo.IsAdmin(userID)
	if err != nil {
		return err
	}
	if !isAdmin {
		return fmt.Errorf("user %d is not an admin", userID)
	}

	return nil
}

func (as *AdminService) ProductExists(productID string) (bool, error) {
	product, err := as.productRepo.GetProductByID(productID)
	if err != nil {
		return false, fmt.Errorf("failed to get product: %v", err)
	}
	return product != nil, nil
}

func (as *AdminService) AddProductService(productID string, name string, stock int, price float64) error {
	exists, err := as.ProductExists(productID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("product with ID %s already exists", productID)
	}

	newProduct := models.Products{
		ProductID: productID,
		Name:      name,
		Stock:     stock,
		Price:     price,
	}

	if err := as.productRepo.AddProduct(&newProduct); err != nil {
		return fmt.Errorf("failed to create new product: %v", err)
	}

	return nil
}

func (as *AdminService) DeleteProductService(productID string) error {
	exists, err := as.ProductExists(productID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("product with ID %s does not exist", productID)
	}

	if err := as.productRepo.DeleteProduct(productID); err != nil {
		return fmt.Errorf("failed to delete product: %v", err)
	}
	return nil
}

func (as *AdminService) UpdateProductService(productID string, stock int, price float64) error {
	exists, err := as.ProductExists(productID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("product with ID %s does not exist", productID)
	}

	productUpdates := models.Products{
		Stock: stock,
		Price: price,
	}

	updatedProduct, err := as.productRepo.UpdateProduct(productID, &productUpdates)
	if _, err := as.productRepo.UpdateProduct(productID, updatedProduct); err != nil {
		return fmt.Errorf("failed to update product: %v", err)
	}

	return nil
}

func validateOrderStatus(status string) error {
	validStatuses := []string{
		string(models.Pending),
		string(models.Confirmed),
		string(models.Preparing),
		string(models.Shipping),
		string(models.Delivered),
	}

	for _, valid := range validStatuses {
		if status == valid {
			return nil
		}
	}
	return fmt.Errorf("invalid order status: %s", status)
}

func validatePaymentStatus(status string) error {
	validStatuses := []string{
		string(models.PaymentPending),
		string(models.PaymentPaid),
	}

	for _, valid := range validStatuses {
		if status == valid {
			return nil
		}
	}
	return fmt.Errorf("invalid payment status: %s", status)
}

func validateShipStatus(status string) error {
	validStatuses := []string{
		string(models.ShipPending),
		string(models.Shipped),
		string(models.Deliver),
	}

	for _, valid := range validStatuses {
		if status == valid {
			return nil
		}
	}
	return fmt.Errorf("invalid shipping status: %s", status)
}

func (as *AdminService) UpdateStatusOrderService(orderID int, orderStatus string, paymentStatus string, shipStatus string) error {
	order, err := as.orderRepo.GetOrderByID(orderID)
	if err != nil {
		return fmt.Errorf("failed to get order: %v", err)
	}
	if order == nil {
		return fmt.Errorf("order with ID %d does not exist", orderID)
	}

	if orderStatus != "" {
		if err := validateOrderStatus(orderStatus); err != nil {
			return fmt.Errorf("invalid order status: %v", err)
		}
		order.OrderStatus = models.OrderStatus(orderStatus)
	}

	if paymentStatus != "" {
		if err := validatePaymentStatus(paymentStatus); err != nil {
			return fmt.Errorf("invalid payment status: %v", err)
		}
		order.PaymentStatus = models.PaymentStatus(paymentStatus)
	}

	if shipStatus != "" {
		if err := validateShipStatus(shipStatus); err != nil {
			return fmt.Errorf("invalid ship status: %v", err)
		}
		order.ShipStatus = models.ShipStatus(shipStatus)
	}

	if err := as.orderRepo.UpdateStatusOrder(order); err != nil {
		return fmt.Errorf("failed to update order: %v", err)
	}
	return nil
}
