package repository

import (
	"errors"
	"tugas_akhir_example/internal/pkg/entity"

	"gorm.io/gorm"
)

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{db: db}
}

func (r *addressRepository) Create(address *entity.Address) error {
	return r.db.Create(address).Error
}

func (r *addressRepository) GetByID(id uint) (*entity.Address, error) {
	var address entity.Address
	if err := r.db.First(&address, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &address, nil
}

func (r *addressRepository) GetByUserID(userID uint, filter entity.FilterAddress) ([]entity.Address, int64, error) {
	var addresses []entity.Address
	var total int64

	query := r.db.Where("user_id = ?", userID)

	if err := query.Model(&entity.Address{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(filter.Limit).Offset(filter.Offset).Find(&addresses).Error; err != nil {
		return nil, 0, err
	}

	return addresses, total, nil
}

func (r *addressRepository) GetByUserIDAndID(userID, addressID uint) (*entity.Address, error) {
	var address entity.Address
	if err := r.db.Where("id = ? AND user_id = ?", addressID, userID).First(&address).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &address, nil
}

func (r *addressRepository) Update(address *entity.Address) error {
	return r.db.Save(address).Error
}

func (r *addressRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Address{}, id).Error
}
