package migrations

import (
	"app/core/config"
	"app/database"
	"app/database/models"
	"fmt"
	"log"
)

func Migrate() {
	config.InitializeAndSet()

	database.InitializeAndConnect()
	db := database.Connection()

	err := db.AutoMigrate(
		&models.User{},
		&models.Token{},
	)

	if err != nil {
		log.Fatal("Cannot migrate")
		return
	}

	fmt.Println("Migration done ..")
}
