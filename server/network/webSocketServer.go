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
var historyConnectedPlayers = make(map[string]string)

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
			player, err := game.MyGame.GetPlayerFromIDSocket(so.Id())
			if err == nil {
				onLeft(player)
			}
		})

		so.On("PUTLINE", func(x int, y int, o int) {
			player, err := game.MyGame.GetPlayerFromIDSocket(so.Id())
			if err == nil {
				onPlayerPlayLine(x, y, o, player.NumPlayer, false)
			}
		})

		so.On("LOGIN", func(user string, pass string) {
			if !accounts.IsValidUsername(user) {
				SendConnectAccept(so, -1, 0) // Refus de la connexion
				return
			}

			// Login (user/pass)
			account, resultatIntLogin := accounts.Login(user, pass)
			if resultatIntLogin == 1 {
				fmt.Println("Connexion réussie pour le client : ", account.Name, " (compte déjà existant)")
			} else if resultatIntLogin == 2 {
				fmt.Println("Connexion réussie pour le client : ", account.Name, " (compte crée)")
			} else if resultatIntLogin == 0 {
				fmt.Println("Tentative de connexion échouée pour le client : ", account.Name, " (mauvais mot de passe)")
			}

			// Entrée dans la partie
			if resultatIntLogin > 0 {
				so.Join("all")
				onPlayerJoin(so, user, resultatIntLogin)
				historyConnectedPlayers[so.Id()] = user
			} else {
				SendConnectAccept(so, resultatIntLogin, 0)
			}

		})

		so.On("GOAGAIN", func(i string) {
			user, _ := historyConnectedPlayers[so.Id()]
			fmt.Println(user + " veut rejouer")
			onPlayerJoin(so, user, 1)
		})

		so.On("ASKLADDER", func(i string) {
			generalLadder := accounts.GetGeneralLadder()
			SendLadder(so, generalLadder)
		})

		// Le client est connecté et est prêt a recevoir les informations
		so.On("READY", func(i string) {
			_, err := game.MyGame.GetPlayerFromIDSocket(so.Id())
			if err == nil {
				onReady(so)
			}
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		globals.ErrLogger.Println("Erreur sur un client : ", err)
	})
}

// sendToAll permet l'envoi d'un paquet à tous les clients
func sendToAll(namePackage string, args ...interface{}) {
	server.BroadcastTo("all", namePackage, args)
	networkLogger.Println("send@all: ", namePackage, args)
}

// sendToClient permet l'envoi d'un paquet à un client
func sendToClient(client socketio.Socket, namePackage string, args ...interface{}) {
	client.Emit(namePackage, args)
	networkLogger.Println("send@client: ", namePackage, args)
}
