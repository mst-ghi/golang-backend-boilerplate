package db_drivers

import (
	"app/core/config"

	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	driver := config.GetDatabaseDriverName()

	var DB *gorm.DB

	if driver == "mysql" {
		DB = MysqlConnect()
	} else if driver == "postgres" {
		DB = PostgresConnect()
	}

	return DB
}
