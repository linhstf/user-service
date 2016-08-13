package product

import "time"

type Category struct {
	ID          int
	Name        string
	ParentID    int
	CreatedDate time.Time
	UpdatedDate time.Time
}
