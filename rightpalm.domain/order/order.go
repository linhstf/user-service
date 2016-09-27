package order

import "time"

// Order determines order information of customer
type Order struct {
	ID string

	CustomerID              string
	CustomerName            string
	CustomerAddress         string
	CustomerSuburb          string
	CustomerCity            string
	CustomerPostcode        string
	CustomerState           string
	CustomerCountry         string
	CustomerPhoneNumber     string
	CustomerEmail           string
	CustomerAddressFormatID int

	DeliveryAddress         string
	DeliverySuburb          string
	DeliveryCity            string
	DeliveryPostcode        string
	DeliveryState           string
	DeliveryCountry         string
	DeliveryPhoneNumber     string
	DeliveryEmail           string
	DeliveryAddressFormatID int

	CCType         string
	CCOwner        string
	CCNumber       string
	CCExpires      string
	CreatedDate    time.Time
	UpdatedDate    time.Time
	PurchasedDate  time.Time
	ShippingCost   float32
	ShippingMethod string
	StatusID       int
	FinishedDate   time.Time
	Comment        string
	Cuurency       string
	CurrencyValye  float32
}
