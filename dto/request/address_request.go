package request

type CreateAddress struct {
	RecipientName        string `json:"recipient_name"`
	RecipientPhoneNumber string `json:"recipient_phone_number"`
	Province             string `json:"province"`
	City                 string `json:"city"`
	District             string `json:"district"`
	Village              string `json:"village"`
	PostalCode           string `json:"postal_code"`
	FullAddress          string `json:"full_address"`
}

type UpdateAddress struct {
	RecipientName        string `json:"recipient_name"`
	RecipientPhoneNumber string `json:"recipient_phone_number"`
	Province             string `json:"province"`
	City                 string `json:"city"`
	District             string `json:"district"`
	Village              string `json:"village"`
	PostalCode           string `json:"postal_code"`
	FullAddress          string `json:"full_address"`
}
