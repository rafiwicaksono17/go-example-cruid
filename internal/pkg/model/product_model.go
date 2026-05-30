package model

type (
	// Product Models
	ProductResponse struct {
		ID          uint    `json:"id"`
		StoreID     uint    `json:"store_id"`
		CategoryID  uint    `json:"category_id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Quantity    int     `json:"quantity"`
		ImageURL    string  `json:"image_url"`
	}

	CreateProductRequest struct {
		CategoryID  uint    `form:"category_id" json:"category_id" validate:"required"`
		Name        string  `form:"name" json:"name" validate:"required"`
		Description string  `form:"description" json:"description"`
		Price       float64 `form:"price" json:"price" validate:"required,gt=0"`
		Quantity    int     `form:"quantity" json:"quantity" validate:"required,gte=0"`
	}

	UpdateProductRequest struct {
		CategoryID  uint    `form:"category_id" json:"category_id"`
		Name        string  `form:"name" json:"name"`
		Description string  `form:"description" json:"description"`
		Price       float64 `form:"price" json:"price"`
		Quantity    int     `form:"quantity" json:"quantity"`
	}

	GetProductsResponse struct {
		Data  []ProductResponse `json:"data"`
		Total int64             `json:"total"`
	}
)
