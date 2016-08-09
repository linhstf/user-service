package customer

import "rightpalm/rightpalm.domain/product"

type CustomerCartAttribute struct {
	ID       int
	Customer Customer
	Product  product.Product
}
