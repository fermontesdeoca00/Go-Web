package domain

// declaration of the struct Product
type Product struct {
	ID           int     `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	Quantity     int     `json:"quantity,omitempty"`
	Code_Value   string  `json:"code_value,omitempty"`
	IS_Published bool    `json:"is_published"`
	Expiration   string  `json:"expiration,omitempty"`
	Price        float64 `json:"price,omitempty"`
}
