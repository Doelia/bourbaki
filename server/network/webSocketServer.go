package network

import (
	"go-bourbaki/server/globals"
	"log"

	"github.com/googollee/go-socket.io"
)

var server *socketio.Server
var err error

func createServerProtocle() {
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Emit("test", "hey")

		so.On("chat", func(msg string) {
			log.Println("Reçu @ chat : ", msg)
			so.BroadcastTo("chat", "chat", msg)
		})

		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})
}

func getWebSocketHandler() *socketio.Server {

	if server, err = socketio.NewServer(nil); err == nil {
		createServerProtocle()
		return server
	}

	globals.ErrLogger.Println("Erreur à la création du protocole : ", err)

	return nil

}
