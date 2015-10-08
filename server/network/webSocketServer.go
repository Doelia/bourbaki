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
		if !game.MyGame.IsPauseNecessary(){
			Unpause()
		}

		so.On("disconnection", func() {
			networkLogger.Println("Un client se déconnecte")
			if game.MyGame.IsPauseNecessary(){
				Pause()
			}
		})

		so.On("PUTLINE", func(x int, y int, o int, n int) { //TODO num joueur déterminé côté serveur
			l := globals.Line{x, y, o, n}
			networkLogger.Println("Un client joue en ", l)
			game.MyGame.AddLine(l)
			DisplayLine(x, y, o, n)
			isSquare, square := game.MyGame.TestSquare(l)
			if isSquare {
				DisplaySquare(square.X, square.Y, square.N) //TODO il faut qu'il rejoue
				if game.MyGame.IsEndGame(){
					//TODO appeller gestionFinPartie
				}
			} else {
				game.MyGame.ChangeCurrentPlayer()
			}
		})

		so.On("LOGIN", func(user string, pass string) {

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
				so.Join("all") // Pour recevoir les broadcasts du game

				var numPlayer int
				player, err := game.MyGame.GetPlayerFromName(user)
				if err != nil { // Pas encore dans la partie
					numPlayer = game.MyGame.GetNewNumPlayer()
					player = game.ConstructPlayer(numPlayer, user)
					game.MyGame.AddPlayer(player)
				} else { // Déjà dans la partie
					numPlayer = player.NumPlayer
				}
				player.IsActive = true
				ConnectAccept(so, resultatIntLogin, numPlayer)

				// On attends le "READY" du joueur pour tout lui envoyer

			} else {
				ConnectAccept(so, resultatIntLogin, 0)
			}

		})

		// Le client est connecté et est pret a recevoir les informations
		so.On("READY", func() {
			// Update de la liste des joueurs pour tout le monde
			UpdatePlayers(game.MyGame.GetAllPlayers())
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
