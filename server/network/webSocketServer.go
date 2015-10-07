package network

import (
	"fmt"
	"go-bourbaki/server/accounts"
	"go-bourbaki/server/game"
	"go-bourbaki/server/globals"

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

		// TODO Passer des string en param
		so.On("PUTLINE", func (x int, y int, o int, n int){ //TODO num joueur déterminé côté serveur
			l := globals.Line{x, y, o, n}
			networkLogger.Println("Un client joue en ", l)
			game.MyGame.AddLine(l)
			DisplayLine(x, y, o, n)
			isSquare, square := game.MyGame.TestSquare(l)
			if isSquare{
				DisplaySquare(square.X, square.Y, square.N)
			}
		})

		so.On("LOGIN", func(user string, pass string) {
			account, resultatIntLogin := accounts.Login(user, pass)
			if resultatIntLogin == 1 {
				fmt.Println("Connexion réussie pour le client : ", account.Name, " (compte déjà existant)")
			} else if resultatIntLogin == 2 {
				fmt.Println("Connexion réussie pour le client : ", account.Name, " (compte crée)")
			} else if resultatIntLogin == 0 {
				fmt.Println("Tentative de connexion échouée pour le client : ", account.Name, " (mauvais mot de passe)")
			}

			if resultatIntLogin > 0 {
				if !game.MyGame.PlayerExists(user) {
					player := game.ConstructPlayer(game.MyGame.GetNewNumPlayer(), user)
					game.MyGame.AddPlayer(player)
				} else {

				}
				ConnectAccept(so, resultatIntLogin, 2)
			} else {
				ConnectAccept(so, resultatIntLogin, 0)
			}

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
