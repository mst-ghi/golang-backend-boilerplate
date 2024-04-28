package core

import (
	"app/pkg/messages"
	"app/pkg/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseTemplate struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Controller struct {
	Service  any
	Response ResponseTemplate
}

type Error map[string]string

type Response[T any] struct {
	Message string
	Errors  struct{ errorKey string }
	Data    T
}

type SuccessResponse struct {
	Message string
	Errors  struct{ errorKey string }
	Data    struct{ dataKey any }
}

var controller *Controller

func GetController() *Controller {
	return controller
}

func ToResponse(message string, data any, errors any) map[string]any {
	return gin.H{
		"message": message,
		"errors":  errors,
		"data":    data,
	}
}

func (ctrl *Controller) Success(c *gin.Context, data any) {
	if data != nil {
		ctrl.Response.Data = data
	}

	response := ToResponse(
		ctrl.Response.Message,
		ctrl.Response.Data,
		map[string]any{},
	)

	c.SecureJSON(ctrl.Response.Status, response)
}

func (ctrl *Controller) JsonBindError(c *gin.Context, err error) {
	response := ToResponse(
		messages.MessageUnprocessable,
		map[string]any{},
		validation.Handle(err),
	)
	c.SecureJSON(http.StatusUnprocessableEntity, response)
}

func (ctrl *Controller) UnprocessableError(c *gin.Context, err Error) {
	response := ToResponse(
		messages.MessageUnprocessable,
		map[string]any{},
		err,
	)
	c.SecureJSON(http.StatusUnprocessableEntity, response)
}

func (ctrl *Controller) NotFoundError(c *gin.Context, err Error) {
	response := ToResponse(
		messages.MessageNotFound,
		map[string]any{},
		err,
	)
	c.SecureJSON(http.StatusNotFound, response)
}

func (ctrl *Controller) BadRequestError(c *gin.Context, err Error) {
	response := ToResponse(
		messages.MessageBadRequest,
		map[string]any{},
		err,
	)
	c.SecureJSON(http.StatusBadRequest, response)
}

func Initialize() {
	controller = &Controller{
		Response: ResponseTemplate{
			Status:  http.StatusOK,
			Message: messages.MessageOk,
			Data:    map[string]any{},
		},
	}
}
