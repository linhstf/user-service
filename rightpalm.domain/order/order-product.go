package order

import "rightpalm/rightpalm.domain/product"

type OrderProduct struct {
	ID         int
	Order      Order
	Product    product.Product
	FinalPrice float32
	Tax        float32
	Quantity   int
}
