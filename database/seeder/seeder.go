package seeder

import (
	"app/core/config"
	"app/database"
)

func Seeder() {
	config.InitializeAndSet()

	database.InitializeAndConnect()
	db := database.Connection()

	UserSeeder(db)
}
