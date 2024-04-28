package config

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var DATABASE_DRIVERS = []string{"mysql", "postgres"}

func InitializeAndSet() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	location, err := time.LoadLocation(Get("APP_TZ"))

	if err != nil {
		log.Fatalf("Err loading location: %v", err)
	}

	time.Local = location
}

func Get(key string) string {
	return os.Getenv(key)
}

func Set(key, value string) error {
	return os.Setenv(key, value)
}

func GetRunAddress() string {
	return Get("SERVER_HOST") + ":" + Get("SERVER_PORT")
}

func GetAppKey() string {
	key := []byte(Get("APP_KEY"))

	if _, err := rand.Read(key); err != nil {
		panic(err.Error())
	}

	return hex.EncodeToString(key)
}

func GetDatabaseDriverName() string {
	driverName := Get("DB_DRIVER")

	if !slices.Contains(DATABASE_DRIVERS, driverName) {
		panic("DB_DRIVER is invalid in .env file, it must be 'mysql' or 'postgres")
	}

	return driverName
}

func GetDatabaseDNS() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		Get("DB_USERNAME"),
		Get("DB_PASSWORD"),
		Get("DB_HOST"),
		Get("DB_PORT"),
		Get("DB_NAME"),
	)
}

func GetTokensExpires() (accessExpiresHours, refreshExpiresDays int) {
	accessExpires, _ := strconv.Atoi(Get("TOKEN_ACCESS_EXPIRES_HOURS"))
	accessExpiresHours = accessExpires

	refreshExpires, _ := strconv.Atoi(Get("TOKEN_REFRESH_EXPIRES_DAYS"))
	refreshExpiresDays = refreshExpires

	return
}
