package entity

import "gorm.io/gorm"

type (
	Category struct {
		ID        uint   `gorm:"primaryKey"`
		Name      string `gorm:"type:varchar(191);not null;uniqueIndex"`
		Products  []Product
		CreatedAt int64
		UpdatedAt int64
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}

	FilterCategory struct {
		Limit  int
		Offset int
		Name   string
	}
)
