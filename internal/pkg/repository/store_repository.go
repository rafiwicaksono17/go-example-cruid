package repository

import (
	"errors"
	"tugas_akhir_example/internal/pkg/entity"

	"gorm.io/gorm"
)

type storeRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) StoreRepository {
	return &storeRepository{db: db}
}

func (r *storeRepository) Create(store *entity.Store) error {
	return r.db.Create(store).Error
}

func (r *storeRepository) GetByID(id uint) (*entity.Store, error) {
	var store entity.Store
	if err := r.db.First(&store, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &store, nil
}

func (r *storeRepository) GetByUserID(userID uint) (*entity.Store, error) {
	var store entity.Store
	if err := r.db.Where("user_id = ?", userID).First(&store).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &store, nil
}

func (r *storeRepository) GetAll(filter entity.FilterStore) ([]entity.Store, int64, error) {
	var stores []entity.Store
	var total int64

	query := r.db

	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}

	if err := query.Model(&entity.Store{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(filter.Limit).Offset(filter.Offset).Find(&stores).Error; err != nil {
		return nil, 0, err
	}

	return stores, total, nil
}

func (r *storeRepository) Update(store *entity.Store) error {
	return r.db.Save(store).Error
}

func (r *storeRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Store{}, id).Error
}
