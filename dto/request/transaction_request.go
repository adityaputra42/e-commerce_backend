package request

type CreateTransaction struct {
	AddressID       int64         `json:"address_id"`
	ShippingID      int64         `json:"shipping_id"`
	PaymentMethodID int64         `json:"payment_method_id"`
	ShippingPrice   float64       `json:"shipping_price"`
	TotalPrice      float64       `json:"total_price"`
	ProductOrders   []CreateOrder `json:"product_orders"`
}

type CreateOrder struct {
	ProductID     int64   `json:"product_id"`
	ColorVarianID int64   `json:"color_varian_id"`
	SizeVarianID  int64   `json:"size_varian_id"`
	UnitPrice     float64 `json:"unit_price"`
	Subtotal      float64 `json:"subtotal"`
	Quantity      int64   `json:"quantity"`
}

type UpdateTransaction struct {
	TxID   string `json:"tx_id"`
	Status string `json:"status"`
}

type UpdateOrder struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}
