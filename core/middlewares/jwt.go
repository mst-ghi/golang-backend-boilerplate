package middlewares

import (
	"app/core"
	"app/database/repositories"
	"app/pkg/messages"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")

	if len(jwtToken) == 0 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[len(jwtToken)-1], nil
}

func Auth(c *gin.Context) {
	accessToken, err := extractBearerToken(c.GetHeader("Authorization"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, core.ToResponse(
			messages.MessageInvalidToken,
			map[string]any{},
			map[string]any{},
		))
		return
	}

	tokenRepo := repositories.NewTokenRepository()
	token := tokenRepo.FindByAccess(accessToken)

	if token.ID == "" || token.User.ID == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, core.ToResponse(
			messages.MessageInvalidToken,
			map[string]any{},
			map[string]any{},
		))
	}

	c.Set(gin.AuthUserKey, token.User)
	c.Set("token", token.AccessToken)

	c.Next()
}
