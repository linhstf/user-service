package customer

import (
	"errors"
	t "time"
)

// Customer store the basic info of customer
type Customer struct {
	ID             string
	Gender         string
	FirstName      string
	LastName       string
	DOB            t.Time
	Email          string
	PhoneNumber    string
	HashedPassword string
}

func NewCustomer(phoneNumber string, email string) (*Customer, error) {
	if phoneNumber == "" {
		return nil, errors.New("")
	}

	newCustomer := new(Customer)
	newCustomer.Email = email
	newCustomer.PhoneNumber = phoneNumber

	return newCustomer, nil
}

func (c *Customer) UpdateCustomer(gender string, firstName string, lastName string, DoB t.Time) *Customer {
	c.Gender = gender
	c.FirstName = firstName
	c.LastName = lastName
	c.DOB = DoB

	return c
}
