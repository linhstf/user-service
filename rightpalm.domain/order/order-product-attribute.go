package order

type OrderProductAttribute struct {
	ID                 int
	OrderID            string
	ProductID          string
	ProductOption      string
	ProductOptionValue string
	OptionValuePrice   float32
	PricePrefix        string
}
