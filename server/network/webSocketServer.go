package network

import (
	"log"

	"github.com/googollee/go-socket.io"
)

var server *socketio.Server
var err error

func getWebSocketHandler() *socketio.Server {

	if server, err = socketio.NewServer(nil); err != nil {
		createServerProtocle()
		return server
	}

	return nil
}

func createServerProtocle() {
	server.On("connection", func(so socketio.Socket) {
		so.Join("game")

		so.Emit("test", "hello :)")

		so.On("chat message", func(msg string) {
			so.BroadcastTo("chat", "chat message", msg)
		})

		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})

}
