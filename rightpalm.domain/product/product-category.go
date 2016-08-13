package product

// ProductCategory indicate that many-many relationship between product and category
type ProductCategory struct {
	ProductID  string
	CategoryID int
}
