package location

type AddressBook struct {
    ID int
    Customer customer.Customer
    Company string
    Street string
    Suburb string
    Postcode string
    City string
    State string
    Country Country
    Zone Zone
}