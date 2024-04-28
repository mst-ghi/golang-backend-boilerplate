package repositories

import (
	"app/database"
	"app/database/models"
	"app/pkg/helpers"
	"sync"
	"time"

	"gorm.io/gorm"
)

type TokenRepositoryInterface interface {
	Connection() *gorm.DB
	Clearing(userId string, wg *sync.WaitGroup)
	Create(userId string) (string, string, string)
	FindByAccess(access string) models.Token
	DeleteByAccess(access string)
	FindByRefresh(refresh string) models.Token
	FindByRefreshAndAccess(access, refresh string) models.Token
}

type TokenRepository struct {
	DB *gorm.DB
}

func NewTokenRepository() *TokenRepository {
	return &TokenRepository{
		DB: database.Connection(),
	}
}

func (repo *TokenRepository) Connection() *gorm.DB {
	return repo.DB
}

func (repo *TokenRepository) Clearing(userId string, wg *sync.WaitGroup) {
	defer wg.Done()
	repo.DB.
		Where("refresh_expires_at < ?", time.Now()).
		Where("user_id = ?", userId).
		Delete(&models.Token{})
}

func (repo *TokenRepository) Create(userId string) (string, string, string) {
	var token models.Token

	accessTokenUUID, _ := helpers.UUID()
	refreshTokenUUID, _ := helpers.UUID()

	token.UserID = userId
	token.AccessToken = accessTokenUUID
	token.RefreshToken = refreshTokenUUID
	token.Invoked = false

	repo.DB.Create(&token).Scan(&token)

	return accessTokenUUID, refreshTokenUUID, token.AccessExpiresAt.Format(time.RFC3339)
}

func (repo *TokenRepository) FindByAccess(access string) models.Token {
	var token models.Token

	repo.DB.
		Joins("User").
		Where(&models.Token{AccessToken: helpers.SHA1(access)}, "access_token", "invoked").
		Where("access_expires_at > ?", time.Now()).
		First(&token)

	return token
}

func (repo *TokenRepository) DeleteByAccess(access string) {
	repo.DB.
		Where(&models.Token{AccessToken: access}, "access_token").
		Delete(&models.Token{})
}

func (repo *TokenRepository) FindByRefresh(refresh string) models.Token {
	var token models.Token
	repo.DB.
		Joins("User").
		Where(&models.Token{RefreshToken: helpers.SHA1(refresh)}, "refresh_token", "invoked").
		Where("refresh_expires_at > ?", time.Now()).
		First(&token)
	return token
}

func (repo *TokenRepository) FindByRefreshAndAccess(access, refresh string) models.Token {
	var token models.Token
	repo.DB.
		Joins("User").
		Where(&models.Token{AccessToken: helpers.SHA1(access), RefreshToken: helpers.SHA1(refresh)}, "access_token", "refresh_token", "invoked").
		Where("refresh_expires_at > ?", time.Now()).
		First(&token)
	return token
}
