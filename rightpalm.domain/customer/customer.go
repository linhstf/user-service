package customer

import t "time"

// Customer store the basic info of customer
type Customer struct {
	ID             int
	Gender         string
	FirstName      string
	LastName       string
	DOB            t.Time
	Email          string
	PhoneNumber    string
	FBAccessToken  string
	HashedPassword string
}
