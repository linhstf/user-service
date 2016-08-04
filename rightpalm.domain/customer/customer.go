package customer

import "time"

// Customer store the basic info of customer
type Customer struct {
	ID          int
	Gender      string
	FirstName   string
	LastName    string
	DOB         time.Time
	Email       string
	PhoneNumber string
	Password    string
}
