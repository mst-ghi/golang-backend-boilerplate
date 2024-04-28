package models

import (
	"app/pkg/helpers"
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"type:varchar(40);primarykey"`
	Name      string `gorm:"type:varchar(191);not null"`
	Email     string `gorm:"type:varchar(191);unique;not null"`
	Password  string `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := HashPassword(u.Password)

	if err != nil {
		log.Fatal("Password hashing failed!")
	}

	u.ID = helpers.CUID()
	u.Password = hashedPassword
	return
}

func (u *User) CheckPasswordHash(password string) bool {
	return helpers.CompareStringAndHash(password, u.Password)
}

func HashPassword(password string) (string, error) {
	return helpers.StringToHash(password)
}
