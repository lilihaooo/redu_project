package massage

import "time"

// Order 订单结构体
type Order struct {
	ID              uint      `gorm:"primarykey" json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	OrderNum        string
	UserID          int
	OrderStatus     string
	PaymentStatus   string
	PaymentMethod   string
	ShippingAddress string
	ShippingMethod  string
	ShippingFee     float64
	TotalAmount     float64
	Note            string
	OrderItems      []OrderItem
}
