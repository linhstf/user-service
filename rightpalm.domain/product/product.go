package product

import "time"

type Product struct {
	ID            int
	Quantity      int
	Model         string
	Image         string
	Price         float32
	AddedDate     time.Time
	UpdatedDate   time.Time
	AvailableDate time.Time
	Weight        int
	Status        bool
	// Tax Tax
	// Manufacture Manufacture
}
