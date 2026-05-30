package entity

import "gorm.io/gorm"

type (
	User struct {
		ID           uint   `gorm:"primaryKey"`
		Email        string `gorm:"type:varchar(191);uniqueIndex;not null"`
		Phone        string `gorm:"type:varchar(191);uniqueIndex;not null"`
		Name         string `gorm:"not null"`
		Password     string `gorm:"not null"`
		IsAdmin      bool   `gorm:"default:false"`
		Store        *Store `gorm:"foreignKey:UserID"`
		Addresses    []Address
		Transactions []Transaction
		CreatedAt    int64
		UpdatedAt    int64
		DeletedAt    gorm.DeletedAt `gorm:"index"`
	}

	FilterUser struct {
		Limit  int
		Offset int
		Name   string
	}
)
