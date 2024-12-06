package response

import "time"

type PaymentMethodResponse struct {
	ID            int64     `json:"id"`
	AccountName   string    `json:"account_name"`
	AccountNumber string    `json:"account_number"`
	BankName      string    `json:"bank_name"`
	BankImages    string    `json:"bank_images"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
}
