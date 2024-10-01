package dto

// Authentication request
type SignupRes struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

type LoginRes struct {
	LoginType string `json:"loginType"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type RefreshTokenRes struct {
	RefreshToken string `json:"refreshToken"`
}

// User request
type UpdateUserRes struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

type CreateOrderRes struct {
	DeliveryAddress string `json:"deliveryAddress"`
	PaymentMethod   string `json:"paymentMethod"`
}

type UpdateOrderRes struct {
	PaymentMethod   string `json:"paymentMethod"`
	DeliveryAddress string `json:"deliveryAddress"`
}

type AddItemRes struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type RemoveItemRes struct {
	ProductId string `json:"productId"`
}

// Admin request
type AddProductRes struct {
	ProductId string  `json:"productId"`
	Name      string  `json:"name"`
	Stock     int     `json:"stock"`
	Price     float64 `json:"price"`
}

type RemoveProductRes struct {
	ProductId string `json:"productId"`
}

type UpdateProductRes struct {
	ProductId string  `json:"productId"`
	Name      string  `json:"name"`
	Stock     int     `json:"stock"`
	Price     float64 `json:"price"`
}

type GetSaleDataRes struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type UpdateOrderStatusRes struct {
	OrderId       uint   `json:"orderId"`
	OrderStatus   string `json:"orderStatus"`
	PaymentStatus string `json:"paymentState"`
	ShipStatus    string `json:"shipStatus"`
}
