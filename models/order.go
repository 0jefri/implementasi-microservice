package models

type Order struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Qty        int     `json:"qty"`
	Price      float64 `json:"price"`
	TotalPrice float64 `json:"total_price"`
	PaymentID  string  `json:"payment_id"`
}
