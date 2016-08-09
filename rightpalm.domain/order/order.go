package order

import (
	"time"

	"rightpalm/rightpalm.domain/customer"

	"rightpalm/rightpalm.domain/location"
)

// Order determines order information of customer
type Order struct {
	ID int

	Customer              customer.Customer
	CustomerName          string
	CustomerAddress       string
	CustomerSuburb        string
	CustomerCity          string
	CustomerPostcode      string
	CustomerState         string
	CustomerCountry       string
	CustomerPhoneNumber   string
	CustomerEmail         string
	CustomerAddressFormat location.AddressFormat

	DeliveryAddress       string
	DeliverySuburb        string
	DeliveryCity          string
	DeliveryPostcode      string
	DeliveryState         string
	DeliveryCountry       string
	DeliveryPhoneNumber   string
	DeliveryEmail         string
	DeliveryAddressFormat location.AddressFormat

	CCType         string
	CCOwner        string
	CCNumber       string
	CCExpires      string
	UpdatedDate    time.Time
	PurchasedDate  time.Time
	ShippingCost   float32
	ShippingMethod string
	Status         string
	FinishedDate   time.Time
	Comment        string
	Cuurency       string
	CurrencyValye  float32
}
