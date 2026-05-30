package mysql

import (
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/entity"

	"gorm.io/gorm"
)

func RunMigration(mysqlDB *gorm.DB) {
	err := mysqlDB.AutoMigrate(
		&entity.User{},
		&entity.Store{},
		&entity.Address{},
		&entity.Category{},
		&entity.Product{},
		&entity.Transaction{},
		&entity.ProductLog{},
	)
	if err != nil {
		helper.Logger(helper.LoggerLevelError, "Failed Database Migrated", err)
	}

	helper.Logger(helper.LoggerLevelInfo, "Database Migrated", nil)
}
