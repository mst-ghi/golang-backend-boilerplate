package gateway

import (
	socketio "github.com/googollee/go-socket.io"
)

func ModuleHandlers() {
	socket.OnEvent("/", EVENT_USER_GET, func(con socketio.Conn) {
		ctx, ok := CheckContext(con)

		if ok {
			SuccessEmitTo(con, EVENT_USER_ME, ctx)
		}
	})
}
