package customer

import "time"

type CustomerCart struct {
	ID         int
	Customer   Customer
	Quantity   int
	FinalPrice float32
	AddedDate  time.Time
}
