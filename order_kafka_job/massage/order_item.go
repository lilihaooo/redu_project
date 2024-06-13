package massage

// OrderItem 订单项结构体
type OrderItem struct {
	ID        int
	OrderID   int
	ProductID int
	Quantity  int
	UnitPrice float64
}
