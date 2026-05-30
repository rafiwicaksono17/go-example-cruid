package entity

import "gorm.io/gorm"

type (
	ProductLog struct {
		ID            uint `gorm:"primaryKey"`
		TransactionID uint `gorm:"not null"`
		Transaction   *Transaction
		ProductID     uint `gorm:"not null"`
		Product       *Product
		Quantity      int
		Price         float64
		CreatedAt     int64
		UpdatedAt     int64
		DeletedAt     gorm.DeletedAt `gorm:"index"`
	}

	FilterProductLog struct {
		Limit         int
		Offset        int
		TransactionID uint
	}
)
