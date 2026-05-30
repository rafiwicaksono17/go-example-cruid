package entity

import "gorm.io/gorm"

type (
	Address struct {
		ID            uint `gorm:"primaryKey"`
		UserID        uint `gorm:"not null"`
		User          *User
		JudulAlamat   string `gorm:"not null"`
		PenerimaNama  string `gorm:"not null"`
		PenerimaPhone string `gorm:"not null"`
		Provinsi      string `gorm:"not null"`
		ProvinsiID    string
		Kabupaten     string `gorm:"not null"`
		KabupatenID   string
		Kecamatan     string `gorm:"not null"`
		KecamatanID   string
		Kelurahan     string `gorm:"not null"`
		KelurahanID   string
		DetailAlamat  string `gorm:"not null"`
		IsDefault     bool   `gorm:"default:false"`
		CreatedAt     int64
		UpdatedAt     int64
		DeletedAt     gorm.DeletedAt `gorm:"index"`
	}

	FilterAddress struct {
		Limit  int
		Offset int
	}
)
