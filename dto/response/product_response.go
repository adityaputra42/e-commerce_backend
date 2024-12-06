package response

import "time"

type ProductResponse struct {
	ID          int64     `json:"id"`
	Category    Category  `json:"category"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Images      string    `json:"images"`
	Rating      float64   `json:"rating"`
	Price       float64   `json:"price"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type ProductDetailResponse struct {
	ID          int64                 `json:"id"`
	Category    Category              `json:"category"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Images      string                `json:"images"`
	Rating      float64               `json:"rating"`
	Price       float64               `json:"price"`
	ColorVarian []ColorVarianResponse `json:"color_varian"`
	UpdatedAt   time.Time             `json:"updated_at"`
	CreatedAt   time.Time             `json:"created_at"`
}

type ColorVarianResponse struct {
	ID         int64                `json:"id"`
	ProductID  int64                `json:"product_id"`
	Name       string               `json:"name"`
	Color      string               `json:"color"`
	Images     string               `json:"images"`
	SizeVarian []SizeVarianResponse `json:"size_varian"`
	UpdatedAt  time.Time            `json:"updated_at"`
	CreatedAt  time.Time            `json:"created_at"`
}

type SizeVarianResponse struct {
	ID            int64     `json:"id"`
	ColorVarianID int64     `json:"color_varian_id"`
	Size          string    `json:"size"`
	Stock         int64     `json:"stock"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
}
