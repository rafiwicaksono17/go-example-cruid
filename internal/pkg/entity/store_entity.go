package entity

import "gorm.io/gorm"

type (
	Store struct {
		ID        uint `gorm:"primaryKey"`
		UserID    uint `gorm:"not null;uniqueIndex"`
		User      *User
		Name      string `gorm:"not null"`
		Address   string
		Phone     string
		Products  []Product
		CreatedAt int64
		UpdatedAt int64
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}

	FilterStore struct {
		Limit  int
		Offset int
		Name   string
	}
)
