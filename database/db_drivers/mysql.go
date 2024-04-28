package db_drivers

import (
	"app/core/config"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MysqlDNS() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Get("DB_USERNAME"),
		config.Get("DB_PASSWORD"),
		config.Get("DB_HOST"),
		config.Get("DB_PORT"),
		config.Get("DB_NAME"),
	)
}

func MysqlConnect() *gorm.DB {
	dialector := mysql.Open(MysqlDNS())

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Error,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      true,
				Colorful:                  true,
			},
		),
	})

	if err != nil {
		panic("Cannot connect to mysql database")
	}

	return db
}
