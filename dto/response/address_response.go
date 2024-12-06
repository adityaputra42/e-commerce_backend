package response

import "time"

type AddressResponse struct {
	ID                   int64     `json:"id"`
	RecipientName        string    `json:"recipient_name"`
	RecipientPhoneNumber string    `json:"recipient_phone_number"`
	Province             string    `json:"province"`
	City                 string    `json:"city"`
	District             string    `json:"district"`
	Village              string    `json:"village"`
	PostalCode           string    `json:"postal_code"`
	FullAddress          string    `json:"full_address"`
	UpdatedAt            time.Time `json:"updated_at"`
	CreatedAt            time.Time `json:"created_at"`
}
