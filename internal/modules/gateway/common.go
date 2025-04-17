package gateway

import (
	"app/pkg/helpers"
)

type CuidGenerator struct {
	ID string
}

func (ig *CuidGenerator) NewID() string {
	return helpers.CUID()
}

type SocketUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SocketContext struct {
	User SocketUser `json:"user"`
}

type SocketData struct {
	Event   string      `json:"event"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}
