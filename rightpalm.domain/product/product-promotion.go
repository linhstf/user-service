package product

import "time"

type ProductPromotion struct {
	ID          int
	ProductID   string
	AppliedDate time.Time
	EndedDate   time.Time
	NewPrice    float64
}
