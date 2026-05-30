package repository

import (
	"errors"
	"tugas_akhir_example/internal/pkg/entity"

	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(transaction *entity.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) GetByID(id uint) (*entity.Transaction, error) {
	var transaction entity.Transaction
	if err := r.db.Preload("ProductLogs").First(&transaction, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &transaction, nil
}

func (r *transactionRepository) GetByUserID(userID uint, filter entity.FilterTransaction) ([]entity.Transaction, int64, error) {
	var transactions []entity.Transaction
	var total int64

	query := r.db.Where("user_id = ?", userID).Preload("ProductLogs")

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	if err := query.Model(&entity.Transaction{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(filter.Limit).Offset(filter.Offset).Order("created_at DESC").Find(&transactions).Error; err != nil {
		return nil, 0, err
	}

	return transactions, total, nil
}

func (r *transactionRepository) GetAll(filter entity.FilterTransaction) ([]entity.Transaction, int64, error) {
	var transactions []entity.Transaction
	var total int64

	query := r.db.Preload("ProductLogs")

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	if err := query.Model(&entity.Transaction{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(filter.Limit).Offset(filter.Offset).Order("created_at DESC").Find(&transactions).Error; err != nil {
		return nil, 0, err
	}

	return transactions, total, nil
}

func (r *transactionRepository) Update(transaction *entity.Transaction) error {
	return r.db.Save(transaction).Error
}

func (r *transactionRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Transaction{}, id).Error
}

func (r *transactionRepository) GetByInvoiceNumber(invoiceNumber string) (*entity.Transaction, error) {
	var transaction entity.Transaction
	if err := r.db.Where("invoice_number = ?", invoiceNumber).Preload("ProductLogs").First(&transaction).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &transaction, nil
}

type productLogRepository struct {
	db *gorm.DB
}

func NewProductLogRepository(db *gorm.DB) ProductLogRepository {
	return &productLogRepository{db: db}
}

func (r *productLogRepository) Create(productLog *entity.ProductLog) error {
	return r.db.Create(productLog).Error
}

func (r *productLogRepository) GetByID(id uint) (*entity.ProductLog, error) {
	var productLog entity.ProductLog
	if err := r.db.First(&productLog, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &productLog, nil
}

func (r *productLogRepository) GetByTransactionID(transactionID uint, filter entity.FilterProductLog) ([]entity.ProductLog, int64, error) {
	var productLogs []entity.ProductLog
	var total int64

	query := r.db.Where("transaction_id = ?", transactionID)
	if err := query.Model(&entity.ProductLog{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(filter.Limit).Offset(filter.Offset).Order("created_at DESC").Find(&productLogs).Error; err != nil {
		return nil, 0, err
	}

	return productLogs, total, nil
}

func (r *productLogRepository) Update(productLog *entity.ProductLog) error {
	return r.db.Save(productLog).Error
}

func (r *productLogRepository) Delete(id uint) error {
	return r.db.Delete(&entity.ProductLog{}, id).Error
}
