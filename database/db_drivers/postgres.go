package db_drivers

import (
	"app/core/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresDNS() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Get("DB_HOST"),
		config.Get("DB_USERNAME"),
		config.Get("DB_PASSWORD"),
		config.Get("DB_NAME"),
		config.Get("DB_PORT"),
		config.Get("DB_SSL"),
		config.Get("APP_TZ"),
	)
}

func PostgresConnect() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  PostgresDNS(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to postgres database")
	}

	return db
}
