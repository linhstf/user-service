package customer

import t "time"

type CustomerCart struct {
	ID         int
	Customer   Customer
	Quantity   int
	FinalPrice float32
	AddedDate  t.Time
}
