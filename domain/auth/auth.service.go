package auth

import (
	"app/core"
	"app/database/models"
	"app/database/repositories"
	"sync"
)

type AuthServiceInterface interface {
	Login(dto LoginDto) (Tokens, core.Error)
	Register(dto RegisterDto) core.Error
	Refresh(dto RefreshDto) (Tokens, core.Error)
	Logout(token string)
	ChangePassword(user models.User, dto PasswordDto) core.Error
}

type AuthService struct {
	userRepository  repositories.UserRepositoryInterface
	tokenRepository repositories.TokenRepositoryInterface
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepository:  repositories.NewUserRepository(),
		tokenRepository: repositories.NewTokenRepository(),
	}
}

func (service *AuthService) Login(dto LoginDto) (Tokens, core.Error) {
	var tokens Tokens

	user := service.userRepository.FindByEmail(dto.Email)

	if user.ID == "" || !user.CheckPasswordHash(dto.Password) {
		return tokens, core.Error{"email": "Invalid input data"}
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go service.tokenRepository.Clearing(user.ID, &waitGroup)

	accessToken, refreshToken, expiresAt := service.tokenRepository.Create(user.ID)

	tokens.AccessToken = accessToken
	tokens.RefreshToken = refreshToken
	tokens.ExpiresAt = expiresAt

	waitGroup.Wait()
	return tokens, nil
}

func (service *AuthService) Register(dto RegisterDto) core.Error {
	user := service.userRepository.FindByEmail(dto.Email)

	if user.ID != "" {
		return core.Error{"email": "User exist with this email"}
	}

	user = models.User{Name: dto.Name, Email: dto.Email, Password: dto.Password}
	service.userRepository.Create(user)

	return nil
}

func (service *AuthService) Refresh(dto RefreshDto) (Tokens, core.Error) {
	var tokens Tokens

	token := service.tokenRepository.FindByRefreshAndAccess(dto.AccessToken, dto.RefreshToken)

	if token.ID == "" {
		return tokens, core.Error{"access_token": "Access token is invalid", "refresh_token": "Refresh token is invalid"}
	}

	user := token.User

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go service.tokenRepository.Clearing(user.ID, &waitGroup)

	accessToken, refreshToken, expiresAt := service.tokenRepository.Create(user.ID)

	tokens.AccessToken = accessToken
	tokens.RefreshToken = refreshToken
	tokens.ExpiresAt = expiresAt

	waitGroup.Wait()

	service.tokenRepository.DeleteByAccess(token.AccessToken)
	return tokens, nil
}

func (service *AuthService) Logout(token string) {
	service.tokenRepository.DeleteByAccess(token)
}

func (service *AuthService) ChangePassword(user models.User, dto PasswordDto) core.Error {
	if user.ID == "" || !user.CheckPasswordHash(dto.CurrentPassword) {
		return core.Error{"current_password": "Current password is invalid"}
	}

	hashedPassword, err := models.HashPassword(dto.NewPassword)

	if err != nil {
		return core.Error{"reason": "Password hashing has error"}
	}

	service.userRepository.Connection().Model(&user).Update("password", hashedPassword)

	return nil
}
