package order

type OrderProduct struct {
	ID          int
	OrderID     string
	ProductID   string
	ProductName string
	FinalPrice  float32
	Tax         float32
	Quantity    int
}
