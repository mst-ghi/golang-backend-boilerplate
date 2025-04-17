package handlers

import (
	"app/internal"
	"app/pkg/messages"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
)

type HttpResponse struct {
	Message     string
	Status      int
	Description string
}

func InternalErrorHandler(c *gin.Context, err any) {
	goErr := errors.Wrap(err, 2)
	c.AbortWithStatusJSON(
		500,
		internal.ToResponse(
			messages.MessageInternalError,
			map[string]any{},
			map[string]any{
				"reason": goErr.Error(),
			},
		),
	)
}
