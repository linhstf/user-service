package product

import "time"

type Product struct {
	ID            int
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
	FBPhotoID     string
	Category      Category
	// Tax Tax
	// Manufacture Manufacture
}

// MakeNew return new product
func MakeNew(fbPhotoID string, name string) *Product {
	if fbPhotoID == "" {
		return nil
	}

	return &Product{
		FBPhotoID: fbPhotoID,
		Name:      name,
	}
}
