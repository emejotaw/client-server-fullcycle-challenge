package database

import (
	"github.com/emejotaw/client-server-fullcycle-challenge/internal/infra/database/entity"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func GetConnection() (*gorm.DB, error) {

	dsn := "database.db"
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
}

func AutoMigrate(db *gorm.DB) error {

	return db.AutoMigrate(&entity.DollarQuotation{})
}
