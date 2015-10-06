package network

import (
	"go-bourbaki/server/globals"
	"go-bourbaki/server/accounts"
	"github.com/googollee/go-socket.io"
)

var server *socketio.Server
var err error

func createWebSocketHandler() *socketio.Server {
	if server, err = socketio.NewServer(nil); err == nil {
		createServerProtocle(server)
		return server
	}
	globals.ErrLogger.Println("Erreur à la création du protocole : ", err)
	return nil
}

func createServerProtocle(*socketio.Server) {
	server.On("connection", func(so socketio.Socket) {
		networkLogger.Println("Un client se connecte")

		so.On("disconnection", func() {
			networkLogger.Println("Un client se déconnecte")
		})

		so.On("LOGIN", func(user string, pass string) {
			account, resultatIntLogin := accounts.Login(user,pass)
			if resultatIntLogin == 1{
				networkLogger.Println("Connexion réussie pour le client : ", account.Name, " (compte déjà existant)")
			} else if resultatIntLogin == 2{
				networkLogger.Println("Connexion réussie pour le client : ", account.Name, " (compte crée)")
			} else if resultatIntLogin == 0{
				networkLogger.Println("Tentative de connexion échouée pour le client : ", account.Name, " (mauvais mot de passe)")
			}
			ConnectAccept(so, resultatIntLogin, 2)
			so.Join("all") // Pour recevoir les broadcasts du game
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		globals.ErrLogger.Println("Erreur sur un client : ", err)
	})
}

func sentToAll(namePackage string, args ...interface{}) {
	server.BroadcastTo("all", namePackage, args)
	networkLogger.Println("send@all: ", namePackage, args)
}

func sendToClient(client socketio.Socket, namePackage string, args ...interface{}) {
	client.Emit(namePackage, args)
	networkLogger.Println("send@client: ", namePackage, args)
}
