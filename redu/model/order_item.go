package model

// OrderItem 订单项结构体
type OrderItem struct {
	OrderNum  string
	ProductID int
	Quantity  int
	UnitPrice float64
}
