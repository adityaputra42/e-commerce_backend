package request

type CreateProduct struct {
	CategoryID  int64                      `json:"category_id"`
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Images      string                     `json:"images"`
	Rating      float32                    `json:"rating"`
	Price       float64                    `json:"price"`
	ColorVarian []CreateColorVarianProduct `json:"color_varians"`
}

type CreateColorVarianProduct struct {
	ProductID int64                     `json:"product_id"`
	Name      string                    `json:"name"`
	Color     string                    `json:"color"`
	Images    string                    `json:"images"`
	Sizes     []CreateSizeVarianProduct `json:"sizes"`
}

type CreateSizeVarianProduct struct {
	ColorVarianID int64  `json:"color_varian_id"`
	Size          string `json:"size"`
	Stock         int64  `json:"stock"`
}
