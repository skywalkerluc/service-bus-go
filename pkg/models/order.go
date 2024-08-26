package models

type Order struct {
	OrderID   string `json:"order_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
