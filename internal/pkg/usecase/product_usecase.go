package usecase

import (
	"errors"
	"time"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/pkg/model"
	"tugas_akhir_example/internal/pkg/repository"
)

type productUseCase struct {
	productRepo  repository.ProductRepository
	storeRepo    repository.StoreRepository
	categoryRepo repository.CategoryRepository
}

func NewProductUseCase(productRepo repository.ProductRepository, storeRepo repository.StoreRepository, categoryRepo repository.CategoryRepository) ProductUseCase {
	return &productUseCase{
		productRepo:  productRepo,
		storeRepo:    storeRepo,
		categoryRepo: categoryRepo,
	}
}

func (u *productUseCase) CreateProduct(userID uint, req model.CreateProductRequest, imagePath string) (*model.ProductResponse, error) {
	// Get user's store
	store, err := u.storeRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	if store == nil {
		return nil, errors.New("store not found for user")
	}

	// Verify category exists
	category, err := u.categoryRepo.GetByID(req.CategoryID)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, errors.New("category not found")
	}

	now := time.Now().Unix()
	product := &entity.Product{
		StoreID:     store.ID,
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
		ImageURL:    imagePath,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := u.productRepo.Create(product); err != nil {
		return nil, err
	}

	return &model.ProductResponse{
		ID:          product.ID,
		StoreID:     product.StoreID,
		CategoryID:  product.CategoryID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		ImageURL:    product.ImageURL,
	}, nil
}

func (u *productUseCase) GetProductByID(id uint) (*model.ProductResponse, error) {
	product, err := u.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, errors.New("product not found")
	}

	return &model.ProductResponse{
		ID:          product.ID,
		StoreID:     product.StoreID,
		CategoryID:  product.CategoryID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		ImageURL:    product.ImageURL,
	}, nil
}

func (u *productUseCase) GetProductsByStore(userID uint, limit, offset int, filter entity.FilterProduct) (*model.GetProductsResponse, error) {
	// Get user's store
	store, err := u.storeRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	if store == nil {
		return nil, errors.New("store not found for user")
	}

	filter.Limit = limit
	filter.Offset = offset
	filter.StoreID = store.ID

	products, total, err := u.productRepo.GetByStoreID(store.ID, filter)
	if err != nil {
		return nil, err
	}

	var responses []model.ProductResponse
	for _, product := range products {
		responses = append(responses, model.ProductResponse{
			ID:          product.ID,
			StoreID:     product.StoreID,
			CategoryID:  product.CategoryID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    product.Quantity,
			ImageURL:    product.ImageURL,
		})
	}

	return &model.GetProductsResponse{
		Data:  responses,
		Total: total,
	}, nil
}

func (u *productUseCase) GetAllProducts(limit, offset int, name string, categoryID, storeID uint, sortBy string) (*model.GetProductsResponse, error) {
	filter := entity.FilterProduct{
		Limit:      limit,
		Offset:     offset,
		Name:       name,
		CategoryID: categoryID,
		StoreID:    storeID,
		SortBy:     sortBy,
	}

	products, total, err := u.productRepo.GetAll(filter)
	if err != nil {
		return nil, err
	}

	var responses []model.ProductResponse
	for _, product := range products {
		responses = append(responses, model.ProductResponse{
			ID:          product.ID,
			StoreID:     product.StoreID,
			CategoryID:  product.CategoryID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    product.Quantity,
			ImageURL:    product.ImageURL,
		})
	}

	return &model.GetProductsResponse{
		Data:  responses,
		Total: total,
	}, nil
}

func (u *productUseCase) UpdateProduct(userID, productID uint, req model.UpdateProductRequest, imagePath string) (*model.ProductResponse, error) {
	// Get user's store
	store, err := u.storeRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	if store == nil {
		return nil, errors.New("store not found for user")
	}

	// Get product and verify it belongs to user's store
	product, err := u.productRepo.GetByIDAndStoreID(productID, store.ID)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, errors.New("product not found or unauthorized")
	}

	if req.CategoryID != 0 {
		// Verify category exists
		category, err := u.categoryRepo.GetByID(req.CategoryID)
		if err != nil {
			return nil, err
		}
		if category == nil {
			return nil, errors.New("category not found")
		}
		product.CategoryID = req.CategoryID
	}

	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Price != 0 {
		product.Price = req.Price
	}
	if req.Quantity != 0 {
		product.Quantity = req.Quantity
	}
	if imagePath != "" {
		product.ImageURL = imagePath
	}

	product.UpdatedAt = time.Now().Unix()

	if err := u.productRepo.Update(product); err != nil {
		return nil, err
	}

	return &model.ProductResponse{
		ID:          product.ID,
		StoreID:     product.StoreID,
		CategoryID:  product.CategoryID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		ImageURL:    product.ImageURL,
	}, nil
}

func (u *productUseCase) DeleteProduct(userID, productID uint) error {
	// Get user's store
	store, err := u.storeRepo.GetByUserID(userID)
	if err != nil {
		return err
	}

	if store == nil {
		return errors.New("store not found for user")
	}

	// Get product and verify it belongs to user's store
	product, err := u.productRepo.GetByIDAndStoreID(productID, store.ID)
	if err != nil {
		return err
	}

	if product == nil {
		return errors.New("product not found or unauthorized")
	}

	return u.productRepo.Delete(productID)
}
