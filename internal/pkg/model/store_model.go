package model

type (
	// Store Models
	StoreResponse struct {
		ID      uint   `json:"id"`
		UserID  uint   `json:"user_id"`
		Name    string `json:"name"`
		Address string `json:"address"`
		Phone   string `json:"phone"`
	}

	CreateStoreRequest struct {
		Name    string `json:"name" validate:"required"`
		Address string `json:"address"`
		Phone   string `json:"phone"`
	}

	UpdateStoreRequest struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Phone   string `json:"phone"`
	}

	GetStoresResponse struct {
		Data  []StoreResponse `json:"data"`
		Total int64           `json:"total"`
	}
)
