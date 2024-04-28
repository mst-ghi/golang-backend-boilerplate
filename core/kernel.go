package core

import (
	"app/database/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) models.User {
	return c.MustGet(gin.AuthUserKey).(models.User)
}

func PaginateQueries(c *gin.Context) (string, int, int) {
	search := c.DefaultQuery("search", "")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	if page <= 0 {
		page = 1
	}

	take, _ := strconv.Atoi(c.DefaultQuery("take", "20"))
	switch {
	case take > 100:
		take = 100
	case take <= 0:
		take = 20
	}

	return search, page, take
}

func Token(c *gin.Context) string {
	return c.MustGet("token").(string)
}
