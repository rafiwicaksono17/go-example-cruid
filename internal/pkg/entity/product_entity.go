package entity

import "gorm.io/gorm"

type (
	Product struct {
		ID          uint `gorm:"primaryKey"`
		StoreID     uint `gorm:"not null"`
		Store       *Store
		CategoryID  uint `gorm:"not null"`
		Category    *Category
		Name        string `gorm:"not null"`
		Description string
		Price       float64 `gorm:"not null"`
		Quantity    int     `gorm:"not null"`
		ImageURL    string
		ProductLogs []ProductLog
		CreatedAt   int64
		UpdatedAt   int64
		DeletedAt   gorm.DeletedAt `gorm:"index"`
	}

	FilterProduct struct {
		Limit      int
		Offset     int
		Name       string
		CategoryID uint
		StoreID    uint
		SortBy     string // "price_asc", "price_desc", "newest"
	}
)
