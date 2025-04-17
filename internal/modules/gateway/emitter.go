package gateway

import socketio "github.com/googollee/go-socket.io"

func successResponse(event string, data interface{}) SocketData {
	return SocketData{
		Event:   event,
		Message: "Successful processing",
		Errors:  struct{}{},
		Data:    data,
	}
}

func errorResponse(event string, errors interface{}) SocketData {
	return SocketData{
		Event:   event,
		Message: "Error in processing",
		Errors:  errors,
		Data:    struct{}{},
	}
}

func BroadcastToGeneral(event string, data interface{}) {
	socket.BroadcastToRoom("", SOCKET_GENERAL_ROOM, event, successResponse(event, data))
}

func BroadcastToRoom(room, event string, data interface{}) {
	socket.BroadcastToRoom("", room, event, successResponse(event, data))
}

func SuccessEmitTo(s socketio.Conn, event string, data interface{}) {
	s.Emit(event, successResponse(event, data))
}

func ErrorEmitTo(s socketio.Conn, event string, errors interface{}) {
	s.Emit(event, errorResponse(event, errors))
}
