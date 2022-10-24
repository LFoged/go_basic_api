package product

type Product struct {
	ProductID    int    `json:"productId,omitempty"`
	ProductName  string `json:"productName,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
	Quantity     int    `json:"quantity,omitempty"`
}
