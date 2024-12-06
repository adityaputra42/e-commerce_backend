package request

type CreateShipping struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	State string  `json:"state"`
}
