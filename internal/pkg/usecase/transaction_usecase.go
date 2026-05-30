package usecase

import (
	"errors"
	"fmt"
	"time"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/pkg/model"
	"tugas_akhir_example/internal/pkg/repository"
)

type transactionUseCase struct {
	transactionRepo repository.TransactionRepository
	productLogRepo  repository.ProductLogRepository
	productRepo     repository.ProductRepository
	addressRepo     repository.AddressRepository
	userRepo        repository.UserRepository
}

func NewTransactionUseCase(
	transactionRepo repository.TransactionRepository,
	productLogRepo repository.ProductLogRepository,
	productRepo repository.ProductRepository,
	addressRepo repository.AddressRepository,
	userRepo repository.UserRepository,
) TransactionUseCase {
	return &transactionUseCase{
		transactionRepo: transactionRepo,
		productLogRepo:  productLogRepo,
		productRepo:     productRepo,
		addressRepo:     addressRepo,
		userRepo:        userRepo,
	}
}

func (u *transactionUseCase) CreateTransaction(userID uint, req model.CreateTransactionRequest) (*model.TransactionResponse, error) {
	// Verify user exists
	user, err := u.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	// Verify address belongs to user
	address, err := u.addressRepo.GetByUserIDAndID(userID, req.AddressID)
	if err != nil {
		return nil, err
	}

	if address == nil {
		return nil, errors.New("address not found")
	}

	// Calculate total price and create product logs
	var totalPrice float64
	var productLogs []entity.ProductLog

	for _, productItem := range req.Products {
		product, err := u.productRepo.GetByID(productItem.ProductID)
		if err != nil {
			return nil, err
		}

		if product == nil {
			return nil, fmt.Errorf("product %d not found", productItem.ProductID)
		}

		// Check product quantity
		if product.Quantity < productItem.Quantity {
			return nil, fmt.Errorf("product %s out of stock. available: %d", product.Name, product.Quantity)
		}

		// Calculate price
		itemPrice := product.Price * float64(productItem.Quantity)
		totalPrice += itemPrice

		// Prepare product log (will be created after transaction)
		productLogs = append(productLogs, entity.ProductLog{
			ProductID: product.ID,
			Quantity:  productItem.Quantity,
			Price:     product.Price,
		})

		// Update product quantity
		product.Quantity -= productItem.Quantity
		product.UpdatedAt = time.Now().Unix()
		if err := u.productRepo.Update(product); err != nil {
			return nil, err
		}
	}

	// Create transaction
	now := time.Now().Unix()
	invoiceNumber := fmt.Sprintf("INV-%d-%d", userID, now)

	transaction := &entity.Transaction{
		UserID:        userID,
		AddressID:     req.AddressID,
		InvoiceNumber: invoiceNumber,
		TotalPrice:    totalPrice,
		Status:        "pending",
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	if err := u.transactionRepo.Create(transaction); err != nil {
		return nil, err
	}

	// Create product logs
	for i := range productLogs {
		productLogs[i].TransactionID = transaction.ID
		productLogs[i].CreatedAt = now
		productLogs[i].UpdatedAt = now
		if err := u.productLogRepo.Create(&productLogs[i]); err != nil {
			return nil, err
		}
	}

	// Fetch transaction with logs
	finalTransaction, err := u.transactionRepo.GetByID(transaction.ID)
	if err != nil {
		return nil, err
	}

	return u.mapTransactionToResponse(finalTransaction), nil
}

func (u *transactionUseCase) GetTransactionByID(userID, transactionID uint) (*model.TransactionResponse, error) {
	transaction, err := u.transactionRepo.GetByID(transactionID)
	if err != nil {
		return nil, err
	}

	if transaction == nil {
		return nil, errors.New("transaction not found")
	}

	// Check authorization
	if transaction.UserID != userID {
		return nil, errors.New("unauthorized: you can only view your own transactions")
	}

	return u.mapTransactionToResponse(transaction), nil
}

func (u *transactionUseCase) GetTransactionsByUser(userID uint, limit, offset int, status string) (*model.GetTransactionsResponse, error) {
	filter := entity.FilterTransaction{
		Limit:  limit,
		Offset: offset,
		Status: status,
	}

	transactions, total, err := u.transactionRepo.GetByUserID(userID, filter)
	if err != nil {
		return nil, err
	}

	var responses []model.TransactionResponse
	for _, transaction := range transactions {
		responses = append(responses, *u.mapTransactionToResponse(&transaction))
	}

	return &model.GetTransactionsResponse{
		Data:  responses,
		Total: total,
	}, nil
}

func (u *transactionUseCase) GetAllTransactions(limit, offset int, status string) (*model.GetTransactionsResponse, error) {
	filter := entity.FilterTransaction{
		Limit:  limit,
		Offset: offset,
		Status: status,
	}

	transactions, total, err := u.transactionRepo.GetAll(filter)
	if err != nil {
		return nil, err
	}

	var responses []model.TransactionResponse
	for _, transaction := range transactions {
		responses = append(responses, *u.mapTransactionToResponse(&transaction))
	}

	return &model.GetTransactionsResponse{
		Data:  responses,
		Total: total,
	}, nil
}

func (u *transactionUseCase) UpdateTransactionStatus(userID, transactionID uint, req model.UpdateTransactionStatusRequest) (*model.TransactionResponse, error) {
	transaction, err := u.transactionRepo.GetByID(transactionID)
	if err != nil {
		return nil, err
	}

	if transaction == nil {
		return nil, errors.New("transaction not found")
	}

	// Check authorization
	if transaction.UserID != userID {
		return nil, errors.New("unauthorized: you can only update your own transactions")
	}

	transaction.Status = req.Status
	transaction.UpdatedAt = time.Now().Unix()

	if err := u.transactionRepo.Update(transaction); err != nil {
		return nil, err
	}

	// Fetch updated transaction
	updatedTransaction, err := u.transactionRepo.GetByID(transaction.ID)
	if err != nil {
		return nil, err
	}

	return u.mapTransactionToResponse(updatedTransaction), nil
}

func (u *transactionUseCase) mapTransactionToResponse(transaction *entity.Transaction) *model.TransactionResponse {
	var productLogs []model.ProductLogResponse
	for _, log := range transaction.ProductLogs {
		productLogs = append(productLogs, model.ProductLogResponse{
			ID:        log.ID,
			ProductID: log.ProductID,
			Quantity:  log.Quantity,
			Price:     log.Price,
		})
	}

	return &model.TransactionResponse{
		ID:            transaction.ID,
		UserID:        transaction.UserID,
		AddressID:     transaction.AddressID,
		InvoiceNumber: transaction.InvoiceNumber,
		TotalPrice:    transaction.TotalPrice,
		Status:        transaction.Status,
		ProductLogs:   productLogs,
	}
}
