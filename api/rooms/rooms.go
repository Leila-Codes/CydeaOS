package rooms

import "github.com/gorilla/websocket"

var (
	messageRooms = make(map[string][]*websocket.Conn)
)

func Broadcast(roomCode, message string) {
	for _, conn := range messageRooms[roomCode] {
		conn.WriteMessage(websocket.TextMessage, []byte(message))
	}
}

func Join(roomCode string, conn *websocket.Conn) {
	messageRooms[roomCode] = append(messageRooms[roomCode], conn)
}

func Leave(roomCode string, conn *websocket.Conn) {
	for i, c := range messageRooms[roomCode] {
		if c == conn {
			messageRooms[roomCode] = append(messageRooms[roomCode][:i], messageRooms[roomCode][i+1:]...)
			return
		}
	}
}
