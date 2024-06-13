package model

import "time"

// Order 订单结构体
type Order struct {
	CreatedAt   time.Time `json:"created_at"`
	OrderNum    string
	UserID      int
	OrderStatus int
	PayStatus   int
	PayMethod   int
	PayAt       *time.Time
	TotalAmount float64
	CompletedAt *time.Time
	CancelledAt *time.Time
	Note        string
	OrderItems  []OrderItem `gorm:"foreignKey:OrderNum;references:OrderNum"`
}
