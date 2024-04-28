package gateway

import (
	"app/database/repositories"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
)

var socket *socketio.Server

func GetSocket() *socketio.Server {
	return socket
}

func Initialize() {
	socket = socketio.NewServer(&engineio.Options{
		SessionIDGenerator: &CuidGenerator{},
	})

	BaseHandlers()
	ModuleHandlers()
}

func Serve(engine *gin.Engine) {
	go func() {
		if err := socket.Serve(); err != nil {
			log.Fatalf("SocketIO listen error: %s\n", err)
		}
	}()

	engine.GET("/socket.io/*any", gin.WrapH(socket))
	engine.POST("/socket.io/*any", gin.WrapH(socket))
}

func BaseHandlers() {
	socket.OnConnect("/", func(con socketio.Conn) error {
		rawToken := getUserTkn(con.URL().RawQuery)

		if rawToken != "" {
			tokenRepo := repositories.NewTokenRepository()
			token := tokenRepo.FindByAccess(rawToken)

			con.SetContext(SocketContext{
				User: SocketUser{
					ID:    token.User.ID,
					Name:  token.User.Name,
					Email: token.User.Email,
				},
			})

			// join client to userID room
			con.Join(token.User.ID)

			// join client to general room
			con.Join(SOCKET_GENERAL_ROOM)

			log.Println("Socket connected by ID:", con.ID(), "Email:", token.User.Email)
		} else {
			con.Close()
		}

		return nil
	})

	socket.OnError("/", func(s socketio.Conn, e error) {
		log.Println("Socket OnError:", e.Error())
	})

	// socket.OnDisconnect("/", func(s socketio.Conn, msg string) {
	// 	log.Println("Socket disconnected by ID:", msg)
	// })
}

func CheckContext(s socketio.Conn) (SocketContext, bool) {
	var ctx, ok = s.Context().(SocketContext)

	if !ok {
		ErrorEmitTo(s, EVENT_ERROR_UNAUTHORIZED, struct{}{})
		return ctx, false
	}

	return ctx, true
}

func getUserTkn(query string) string {
	tkn := strings.Split(query, "&")

	if tkn[0] != "" {
		for i := 0; i < len(tkn); i++ {
			if strings.Contains(tkn[i], "tkn=") {
				return strings.Split(tkn[i], "=")[1]
			}
		}
	}

	return ""
}
