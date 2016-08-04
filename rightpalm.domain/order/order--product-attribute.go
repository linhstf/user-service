package order

import "rightpalm.domain/product"

type OrderProductAttribute struct {
	ID                 int
	Order              Order
	Product            product.Product
	ProductOption      string
	ProductOptionValue string
	OptionValuePrice   float32
	PricePrefix        string
}
