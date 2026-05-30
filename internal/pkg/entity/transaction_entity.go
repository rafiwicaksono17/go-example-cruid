package entity

import "gorm.io/gorm"

type (
	Transaction struct {
		ID            uint `gorm:"primaryKey"`
		UserID        uint `gorm:"not null"`
		User          *User
		AddressID     uint `gorm:"not null"`
		Address       *Address
		InvoiceNumber string `gorm:"type:varchar(191);uniqueIndex;not null"`
		TotalPrice    float64
		Status        string `gorm:"not null;default:'pending'"` // pending, completed, cancelled
		ProductLogs   []ProductLog
		CreatedAt     int64
		UpdatedAt     int64
		DeletedAt     gorm.DeletedAt `gorm:"index"`
	}

	FilterTransaction struct {
		Limit  int
		Offset int
		Status string
	}
)
