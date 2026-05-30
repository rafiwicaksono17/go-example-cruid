package model

type (
	// Transaction Models
	TransactionResponse struct {
		ID            uint                 `json:"id"`
		UserID        uint                 `json:"user_id"`
		AddressID     uint                 `json:"address_id"`
		InvoiceNumber string               `json:"invoice_number"`
		TotalPrice    float64              `json:"total_price"`
		Status        string               `json:"status"`
		ProductLogs   []ProductLogResponse `json:"product_logs"`
	}

	CreateTransactionRequest struct {
		AddressID uint                       `json:"address_id" validate:"required"`
		Products  []CreateTransactionProduct `json:"products" validate:"required,min=1"`
	}

	CreateTransactionProduct struct {
		ProductID uint `json:"product_id" validate:"required"`
		Quantity  int  `json:"quantity" validate:"required,gt=0"`
	}

	UpdateTransactionStatusRequest struct {
		Status string `json:"status" validate:"required,oneof=pending completed cancelled"`
	}

	ProductLogResponse struct {
		ID        uint    `json:"id"`
		ProductID uint    `json:"product_id"`
		Quantity  int     `json:"quantity"`
		Price     float64 `json:"price"`
	}

	GetTransactionsResponse struct {
		Data  []TransactionResponse `json:"data"`
		Total int64                 `json:"total"`
	}
)
