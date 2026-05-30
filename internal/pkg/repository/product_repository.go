package repository

import (
	"errors"
	"tugas_akhir_example/internal/pkg/entity"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *entity.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) GetByID(id uint) (*entity.Product, error) {
	var product entity.Product
	if err := r.db.Preload("ProductLogs").First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) GetByStoreID(storeID uint, filter entity.FilterProduct) ([]entity.Product, int64, error) {
	var products []entity.Product
	var total int64

	query := r.db.Where("store_id = ?", storeID)

	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}
	if filter.CategoryID != 0 {
		query = query.Where("category_id = ?", filter.CategoryID)
	}

	if err := query.Model(&entity.Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Sorting
	if filter.SortBy == "price_asc" {
		query = query.Order("price ASC")
	} else if filter.SortBy == "price_desc" {
		query = query.Order("price DESC")
	} else {
		query = query.Order("created_at DESC")
	}

	if err := query.Limit(filter.Limit).Offset(filter.Offset).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (r *productRepository) GetAll(filter entity.FilterProduct) ([]entity.Product, int64, error) {
	var products []entity.Product
	var total int64

	query := r.db

	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}
	if filter.CategoryID != 0 {
		query = query.Where("category_id = ?", filter.CategoryID)
	}
	if filter.StoreID != 0 {
		query = query.Where("store_id = ?", filter.StoreID)
	}

	if err := query.Model(&entity.Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Sorting
	if filter.SortBy == "price_asc" {
		query = query.Order("price ASC")
	} else if filter.SortBy == "price_desc" {
		query = query.Order("price DESC")
	} else {
		query = query.Order("created_at DESC")
	}

	if err := query.Limit(filter.Limit).Offset(filter.Offset).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (r *productRepository) Update(product *entity.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Product{}, id).Error
}

func (r *productRepository) GetByIDAndStoreID(productID, storeID uint) (*entity.Product, error) {
	var product entity.Product
	if err := r.db.Where("id = ? AND store_id = ?", productID, storeID).Preload("ProductLogs").First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}
