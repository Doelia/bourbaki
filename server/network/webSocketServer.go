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
			player, err := game.MyGame.GetPlayerFromIDSocket(so.Id())
			if err == nil {
				player.IsActive = false
			}
			UpdatePlayers(game.MyGame.GetAllPlayers())
			if game.MyGame.IsPauseNecessary() {
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
				DisplaySquare(square.X, square.Y, square.N)
				if game.MyGame.IsEndGame() {
					//TODO appeller gestionFinPartie
				}
				lastPlayer, _ := game.MyGame.GetPreviousPlayer()
				lastPlayer.Score = lastPlayer.Score - 1
				currentPlayer, _ := game.MyGame.GetPlayerFromNumPlayer(game.MyGame.CurrentPlayer.NumPlayer)
				currentPlayer.Score = currentPlayer.Score + 1
				fmt.Println("Attribution de points")
				UpdatePlayers(game.MyGame.GetAllPlayers())
			} else {
				game.MyGame.ChangeCurrentPlayer()
				SetActivePlayers(game.MyGame.CurrentPlayer.NumPlayer)
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
				so.Join("all") // Pour recevoir les broadcasts du gam<e

				var numPlayer int
				player, err := game.MyGame.GetPlayerFromName(user)
				if err != nil { // Pas encore dans la partie
					numPlayer = game.MyGame.GetNewNumPlayer()
					newPlayer := game.ConstructPlayer(numPlayer, user, so.Id())
					game.MyGame.AddPlayer(newPlayer)
					player, _ = game.MyGame.GetPlayerFromName(user)
				} else { // Déjà dans la partie
					numPlayer = player.NumPlayer
				}
				if game.MyGame.CurrentPlayer.NumPlayer == 0 {
					game.MyGame.CurrentPlayer = *player
				}
				player.IsActive = true
				ConnectAccept(so, resultatIntLogin, numPlayer)

				// On attends le "READY" du joueur pour tout lui envoyer

			} else {
				ConnectAccept(so, resultatIntLogin, 0)
			}

		})

		// Le client est connecté et est pret a recevoir les informations
		so.On("READY", func(i string) {
			UpdatePlayers(game.MyGame.GetAllPlayers())
			SetActivePlayers(game.MyGame.CurrentPlayer.NumPlayer)
			if !game.MyGame.IsPauseNecessary() {
				Unpause()
			} else {
				Pause()
			}
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
