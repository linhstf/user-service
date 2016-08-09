package rightpalmServices

import (
	p "rightpalm/rightpalm.domain/product"
)

func GetProductById(id int) *p.Product {
	if id < 0 {
		return nil
	}

	return p.GetProductById(id)
}
