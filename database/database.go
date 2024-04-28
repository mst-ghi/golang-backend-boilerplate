package database

import (
	"app/database/db_drivers"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() *gorm.DB {
	return DB
}

func InitializeAndConnect() {
	DB = db_drivers.GetDB()
}
