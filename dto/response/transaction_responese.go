package response

import "time"

type TransactionResponse struct {
	TxID          string                `json:"tx_id"`
	Address       AddressResponse       `json:"address"`
	Shipping      ShippingResponse      `json:"shipping"`
	PaymentMethod PaymentMethodResponse `json:"payment_method"`
	ShippingPrice float64               `json:"shipping_price"`
	TotalPrice    float64               `json:"total_price"`
	Status        string                `json:"status"`
	Orders        []OrderResponse       `json:"orders"`
	UpdatedAt     time.Time             `json:"updated_at"`
	CreatedAt     time.Time             `json:"created_at"`
}

type OrderResponse struct {
	ID            string               `json:"id"`
	TransactionID string               `json:"transaction_id"`
	Product       ProductOrderResponse `json:"product"`
	Size          string               `json:"size"`
	UnitPrice     float64              `json:"unit_price"`
	Subtotal      float64              `json:"subtotal"`
	Quantity      int64                `json:"quantity"`
	Status        string               `json:"status"`
	UpdatedAt     time.Time            `json:"updated_at"`
	CreatedAt     time.Time            `json:"created_at"`
}

type ProductOrderResponse struct {
	ID          int64                    `json:"id"`
	Name        string                   `json:"name"`
	Category    Category                 `json:"category"`
	Description string                   `json:"description"`
	Images      string                   `json:"images"`
	Rating      float64                  `json:"rating"`
	Price       float64                  `json:"price"`
	ColorVarian ColorVarianOrderResponse `json:"color_varian"`
	UpdatedAt   time.Time                `json:"updated_at"`
	CreatedAt   time.Time                `json:"created_at"`
}

type ColorVarianOrderResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	Images    string    `json:"images"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
