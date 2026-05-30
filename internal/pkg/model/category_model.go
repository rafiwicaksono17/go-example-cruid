package model

type (
	// Category Models
	CategoryResponse struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	CreateCategoryRequest struct {
		Name string `json:"name" validate:"required"`
	}

	UpdateCategoryRequest struct {
		Name string `json:"name" validate:"required"`
	}

	GetCategoriesResponse struct {
		Data  []CategoryResponse `json:"data"`
		Total int64              `json:"total"`
	}
)
