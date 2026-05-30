package repository

import (
	"tugas_akhir_example/internal/pkg/entity"
)

type (
	UserRepository interface {
		Create(user *entity.User) error
		GetByID(id uint) (*entity.User, error)
		GetByEmail(email string) (*entity.User, error)
		GetByPhone(phone string) (*entity.User, error)
		GetAll(filter entity.FilterUser) ([]entity.User, int64, error)
		Update(user *entity.User) error
		Delete(id uint) error
	}

	StoreRepository interface {
		Create(store *entity.Store) error
		GetByID(id uint) (*entity.Store, error)
		GetByUserID(userID uint) (*entity.Store, error)
		GetAll(filter entity.FilterStore) ([]entity.Store, int64, error)
		Update(store *entity.Store) error
		Delete(id uint) error
	}

	AddressRepository interface {
		Create(address *entity.Address) error
		GetByID(id uint) (*entity.Address, error)
		GetByUserID(userID uint, filter entity.FilterAddress) ([]entity.Address, int64, error)
		GetByUserIDAndID(userID, addressID uint) (*entity.Address, error)
		Update(address *entity.Address) error
		Delete(id uint) error
	}

	CategoryRepository interface {
		Create(category *entity.Category) error
		GetByID(id uint) (*entity.Category, error)
		GetByName(name string) (*entity.Category, error)
		GetAll(filter entity.FilterCategory) ([]entity.Category, int64, error)
		Update(category *entity.Category) error
		Delete(id uint) error
	}

	ProductRepository interface {
		Create(product *entity.Product) error
		GetByID(id uint) (*entity.Product, error)
		GetByStoreID(storeID uint, filter entity.FilterProduct) ([]entity.Product, int64, error)
		GetAll(filter entity.FilterProduct) ([]entity.Product, int64, error)
		Update(product *entity.Product) error
		Delete(id uint) error
		GetByIDAndStoreID(productID, storeID uint) (*entity.Product, error)
	}

	TransactionRepository interface {
		Create(transaction *entity.Transaction) error
		GetByID(id uint) (*entity.Transaction, error)
		GetByUserID(userID uint, filter entity.FilterTransaction) ([]entity.Transaction, int64, error)
		GetAll(filter entity.FilterTransaction) ([]entity.Transaction, int64, error)
		Update(transaction *entity.Transaction) error
		Delete(id uint) error
		GetByInvoiceNumber(invoiceNumber string) (*entity.Transaction, error)
	}

	ProductLogRepository interface {
		Create(productLog *entity.ProductLog) error
		GetByID(id uint) (*entity.ProductLog, error)
		GetByTransactionID(transactionID uint, filter entity.FilterProductLog) ([]entity.ProductLog, int64, error)
		Update(productLog *entity.ProductLog) error
		Delete(id uint) error
	}
)
