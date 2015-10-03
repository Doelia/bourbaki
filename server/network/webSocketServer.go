package network

import (
	"fmt"
	"go-bourbaki/server/globals"
	"log"

	"github.com/googollee/go-socket.io"
)

var server *socketio.Server
var err error

func sentToAll(namePackage string, args ...interface{}) {
	server.BroadcastTo("all", namePackage, args)
	fmt.Println("Broadcast ", namePackage, args)
}

func sendToClient(client socketio.Socket, namePackage string, args ...interface{}) {
	client.Emit(namePackage, args)
}

func createServerProtocle() {
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")

		so.On("disconnection", func() {
			log.Println("on disconnect")
		})

		so.On("LOGIN", func(user string, pass string) {
			// TODO Code à isoler et login à implémenter
			ConnectAccept(so, 1, 2)
			so.Join("all") // Pour recevoir les broadcasts du gane
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
