package customer

import "time"

type CustomerLoginInfo struct {
	ID            int
	LastLogin     time.Time
	NumberOfLogin int
	CreatedDate   time.Time
	UpdatedDate   time.Time
}
