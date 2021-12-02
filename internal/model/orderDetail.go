package model

type OrderDetail struct {
	ID int64 `json:"id"`
	OrderID int64 `json:"order_id"`
	ProductID int64 `json:"product_id"`
	Price float64 `json:"price"`
	Quantity int64 `json:"quantity"`
}

type OrderDetails struct {
	OrderDetails []OrderDetail `json:"order_details"`
}