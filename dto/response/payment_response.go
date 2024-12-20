package response

import (
	"time"
)

type PaymentAdminResponse struct {
	ID            int64     `json:"id"`
	TransactionId string    `json:"transaction_id"`
	TotalPayment  float64   `json:"total_payment"`
	Status        string    `json:"status"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
}

type PaymentResponse struct {
	ID           int64              `json:"id"`
	Transaction  TransactionPayment `json:"transaction"`
	TotalPayment float64            `json:"total_payment"`
	Status       string             `json:"status"`
	UpdatedAt    time.Time          `json:"updated_at"`
	CreatedAt    time.Time          `json:"created_at"`
}

type TransactionPayment struct {
	TxID            string    `json:"tx_id"`
	AddressID       int64     `json:"address_id"`
	ShippingID      int64     `json:"shipping_id"`
	PaymentMethodID int64     `json:"payment_method_id"`
	ShippingPrice   float64   `json:"shipping_price"`
	TotalPrice      float64   `json:"total_price"`
	Status          string    `json:"status"`
	UpdatedAt       time.Time `json:"updated_at"`
	CreatedAt       time.Time `json:"created_at"`
}
