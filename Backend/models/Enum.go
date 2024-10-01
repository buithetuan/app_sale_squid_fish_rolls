package models

type OrderStatus string

const (
	Pending   OrderStatus = "pending"
	Confirmed OrderStatus = "confirmed"
	Preparing OrderStatus = "preparing"
	Shipping  OrderStatus = "shipping"
	Delivered OrderStatus = "delivered"
)

type PaymentStatus string

const (
	PaymentPending PaymentStatus = "pending"
	PaymentPaid    PaymentStatus = "paid"
)

type PaymentMethod string

const (
	UnPaid  PaymentMethod = "unpaid"
	Tranfer PaymentMethod = "transfer"
	Cash    PaymentMethod = "cash"
)

type ShipStatus string

const (
	ShipPending ShipStatus = "pending"
	Deliver     ShipStatus = "deliver"
	Shipped     ShipStatus = "shipped"
)

type UserRank string

const (
	Bronze  UserRank = "Bronze"
	Silver  UserRank = "Silver"
	Gold    UserRank = "Gold"
	Premium UserRank = "Premium"
	Patron  UserRank = "Patron"
)

func GetDiscountRate(rank UserRank) float64 {
	switch rank {
	case Silver:
		return 0.05
	case Gold:
		return 0.08
	case Premium:
		return 0.1
	case Patron:
		return 0.15
	case Bronze:
		return 0.0
	default:
		return 0.0
	}
}
