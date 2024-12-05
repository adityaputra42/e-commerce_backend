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
