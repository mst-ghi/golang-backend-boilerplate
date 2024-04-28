package models

import (
	"app/core/config"
	"app/pkg/helpers"
	"time"

	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type Token struct {
	ID               string    `gorm:"type:varchar(40);primarykey"`
	UserID           string    `gorm:"type:varchar(40);not null"`
	AccessToken      string    `gorm:"type:varchar(64);unique;not null"`
	RefreshToken     string    `gorm:"type:varchar(64);unique;not null"`
	Invoked          bool      `gorm:"default:false"`
	AccessExpiresAt  time.Time `gorm:"not null"`
	RefreshExpiresAt time.Time `gorm:"not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time

	User User `gorm:"foreignkey:UserID"`
}

func (t *Token) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = helpers.CUID()

	t.AccessToken = helpers.SHA1(t.AccessToken)
	t.RefreshToken = helpers.SHA1(t.RefreshToken)

	accessExpiresAt, refreshExpiresAt := GetTokenAccessExpires()

	t.AccessExpiresAt = accessExpiresAt
	t.RefreshExpiresAt = refreshExpiresAt

	return
}

func GetTokenAccessExpires() (accessExpiresAt, refreshExpiresAt time.Time) {
	access, refresh := config.GetTokensExpires()

	accessExpiresAt = carbon.Now().AddHours(access).StdTime()
	refreshExpiresAt = carbon.Now().AddDays(refresh).StdTime()

	return
}
