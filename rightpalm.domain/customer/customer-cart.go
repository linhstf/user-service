package customer

import t "time"

type CustomerCart struct {
	ID          int
	CustomerID  string
	Quantity    int
	FinalPrice  float32
	CreatedDate t.Time
}
