package usecase

import (
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/pkg/model"
)

type (
	AuthUseCase interface {
		Register(req model.RegisterRequest) (*model.AuthResponse, error)
		Login(req model.LoginRequest) (*model.AuthResponse, error)
	}

	UserUseCase interface {
		GetUserByID(id uint) (*model.UserResponse, error)
		GetAllUsers(limit, offset int, name string) (*model.GetUsersResponse, error)
		UpdateUser(id uint, req model.UpdateUserRequest) (*model.UserResponse, error)
		DeleteUser(id uint) error
	}

	StoreUseCase interface {
		CreateStore(userID uint, req model.CreateStoreRequest) (*model.StoreResponse, error)
		GetStoreByUserID(userID uint) (*model.StoreResponse, error)
		GetStoreByID(id uint) (*model.StoreResponse, error)
		GetAllStores(limit, offset int, name string) (*model.GetStoresResponse, error)
		UpdateStore(userID, storeID uint, req model.UpdateStoreRequest) (*model.StoreResponse, error)
		DeleteStore(userID, storeID uint) error
	}

	AddressUseCase interface {
		CreateAddress(userID uint, req model.CreateAddressRequest) (*model.AddressResponse, error)
		GetAddressByID(userID, addressID uint) (*model.AddressResponse, error)
		GetAllAddressesByUser(userID uint, limit, offset int) (*model.GetAddressesResponse, error)
		UpdateAddress(userID, addressID uint, req model.UpdateAddressRequest) (*model.AddressResponse, error)
		DeleteAddress(userID, addressID uint) error
	}

	CategoryUseCase interface {
		CreateCategory(req model.CreateCategoryRequest) (*model.CategoryResponse, error)
		GetCategoryByID(id uint) (*model.CategoryResponse, error)
		GetAllCategories(limit, offset int, name string) (*model.GetCategoriesResponse, error)
		UpdateCategory(id uint, req model.UpdateCategoryRequest) (*model.CategoryResponse, error)
		DeleteCategory(id uint) error
	}

	ProductUseCase interface {
		CreateProduct(userID uint, req model.CreateProductRequest, imagePath string) (*model.ProductResponse, error)
		GetProductByID(id uint) (*model.ProductResponse, error)
		GetProductsByStore(userID uint, limit, offset int, filter entity.FilterProduct) (*model.GetProductsResponse, error)
		GetAllProducts(limit, offset int, name string, categoryID, storeID uint, sortBy string) (*model.GetProductsResponse, error)
		UpdateProduct(userID, productID uint, req model.UpdateProductRequest, imagePath string) (*model.ProductResponse, error)
		DeleteProduct(userID, productID uint) error
	}

	TransactionUseCase interface {
		CreateTransaction(userID uint, req model.CreateTransactionRequest) (*model.TransactionResponse, error)
		GetTransactionByID(userID, transactionID uint) (*model.TransactionResponse, error)
		GetTransactionsByUser(userID uint, limit, offset int, status string) (*model.GetTransactionsResponse, error)
		GetAllTransactions(limit, offset int, status string) (*model.GetTransactionsResponse, error)
		UpdateTransactionStatus(userID, transactionID uint, req model.UpdateTransactionStatusRequest) (*model.TransactionResponse, error)
	}
)
