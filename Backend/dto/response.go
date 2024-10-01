package dto

import (
	"github.com/gin-gonic/gin"
)

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserInfoResponse struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

type CartItemsResponse struct {
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Price       string `json:"price"`
	TotalPrice  string `json:"total_price"`
}

type GetCartItemResponse struct {
	CartID        uint                `json:"cart_id"`
	UserID        uint                `json:"user_id"`
	CartItems     []CartItemsResponse `json:"cart_items"`
	TotalQuantity int                 `json:"total_quantity"`
	TotalPrice    string              `json:"total_price"`
}

type OrderResponse struct {
	OrderID         uint                `json:"order_id"`
	UserID          uint                `json:"user_id"`
	OrderItems      []CartItemsResponse `json:"order_items"`
	TotalPrice      float64             `json:"total_price"`
	ShipPrice       float64             `json:"ship_price"`
	DiscountPrice   float64             `json:"discount_price"`
	DeliveryAddress string              `json:"address"`
	PaymentMethod   string              `json:"payment_method"`
}

type APIResponse struct {
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
}

func SendResponse(c *gin.Context, data interface{}, status int, message string) {
	response := APIResponse{data, status, message}
	c.JSON(status, response)
}
