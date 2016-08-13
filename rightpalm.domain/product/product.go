package product

import "time"

type Product struct {
	ID            string
	Name          string
	Quantity      int
	Model         string
	Image         string
	Price         float32
	CreatedDate   time.Time
	UpdatedDate   time.Time
	AvailableDate time.Time
	Weight        int
	Status        bool
	CategoryID    int
	// Tax Tax
	// Manufacture Manufacture
}
