package helper

const (
	authorizationHeaderKey  = "Authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func GetHeaderKey() string {
	return authorizationHeaderKey
}
func GetTypeBearer() string {
	return authorizationTypeBearer
}
func GetPayloadKey() string {
	return authorizationPayloadKey
}

const (
	WaitingPayment       = "waiting_payment"
	WaitingConfirPayment = "waiting_confirm_payment"
	Confirmed            = "confirmed"
)

const (
	Pending    = "pending"
	Packaged   = "packaged"
	OnDelivery = "on_delivery"
	Delivered  = "delivered"
	Received   = "received"
	Cancelled  = "cancelled"
)

func ValidateStatusOrder(item string) bool {
	listStatus := []string{OnDelivery, Delivered, Received, Cancelled}
	for _, v := range listStatus {
		if v == item {
			return true
		}
	}
	return false
}
