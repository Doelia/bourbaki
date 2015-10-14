package network

import (
	"go-bourbaki/server/game"
	"go-bourbaki/server/globals"
	"log"
	"os"

	"github.com/googollee/go-socket.io"
)

var controllerLogger = log.New(os.Stdout, "[event] ", 0)

// Timer ..
var MyTimer *Timer

// OnCreateGame TODO commentaire
func OnCreateGame() {
	MyTimer = createTimer()
}

func onResume() {
	controllerLogger.Println("onResume()")
	SendUnpause()
	onNewTurn()
}

func onPause() {
	controllerLogger.Println("onPause()")
	SendPause()
}

func onPlayerJoin(so socketio.Socket, user string, resultatIntLogin int) {
	controllerLogger.Println("onPlayerJoin: ", user)

	var numPlayer int
	player, err := game.MyGame.GetPlayerFromName(user)
	if err != nil { // Pas encore dans la partie
		numPlayer = game.MyGame.GetNewNumPlayer()
		newPlayer := game.ConstructPlayer(numPlayer, user, so.Id())
		game.MyGame.AddPlayer(newPlayer)
		player, _ = game.MyGame.GetPlayerFromName(user)
	} else { // Déjà dans la partie
		numPlayer = player.NumPlayer
		player.IDSocket = so.Id()
	}
	if game.MyGame.CurrentPlayer == nil {
		game.MyGame.CurrentPlayer = player
	}
	player.IsActive = true
	SendConnectAccept(so, resultatIntLogin, numPlayer)

	// Suivera d'un onReady()
}

func onReady(so socketio.Socket) {
	controllerLogger.Println("onReady")

	SendGrid(so, game.MyGame.GetActivesLinesList(), game.MyGame.GetActivesSquaresList())
	SendUpdatePlayers(game.MyGame.GetAllPlayers())
	SendSetActivePlayers(game.MyGame.CurrentPlayer.NumPlayer)

	if !game.MyGame.IsPauseNecessary() {
		onResume()
	} else {
		onPause()
	}
}

func onLeft(player *globals.Player) {
	controllerLogger.Println("onLeft: ", player.Name)

	player.IsActive = false
	SendUpdatePlayers(game.MyGame.GetAllPlayers())

	if game.MyGame.IsPauseNecessary() {
		onPause()
	} else {
		if player.NumPlayer == game.MyGame.CurrentPlayer.NumPlayer {
			AI()
		}
	}

}

func onSquareDone(squareStruct globals.Square) {
	controllerLogger.Println("onSquareDone: ", squareStruct)
	game.MyGame.AddSquare(squareStruct)
	SendDisplaySquare(squareStruct.X, squareStruct.Y, squareStruct.N)

	// Gestion des scores
	lastPlayer, _ := game.MyGame.GetPreviousPlayer()
	lastPlayer.Score = lastPlayer.Score - 1
	currentPlayer, _ := game.MyGame.GetPlayerFromNumPlayer(game.MyGame.CurrentPlayer.NumPlayer)
	currentPlayer.Score = currentPlayer.Score + 1

	SendUpdatePlayers(game.MyGame.GetAllPlayers())
}

func onNewTurn() {
	controllerLogger.Println("onNewTurn: ", game.MyGame.CurrentPlayer)

	MyTimer.Cancel()

	if !game.MyGame.CurrentPlayer.IsActive {
		AI()
	} else {
		MyTimer.LaunchNewTimer()
		SendSetActivePlayers(game.MyGame.CurrentPlayer.NumPlayer)
	}
}

func onPlayerPlayLine(x int, y int, o int, n int) {

	if game.MyGame.IsPauseNecessary() {
		return
	}

	l := globals.Line{x, y, o, n}
	controllerLogger.Println("onPlayerPlayLine: ", game.MyGame.CurrentPlayer.Name, l)

	game.MyGame.AddLine(l)
	SendDisplayLine(x, y, o, n)

	isSquare, squares := game.MyGame.TestSquare(l)
	if isSquare {
		for _, squareStruct := range squares {
			onSquareDone(squareStruct)
		}
	} else {
		game.MyGame.ChangeCurrentPlayer()
	}
	game.MyGame.GetLadder()
	if game.MyGame.IsEndGame() {
		onEndGame()
	} else {
		onNewTurn()
	}
}

// AI ...
func AI() {
	controllerLogger.Println("AI pour " + game.MyGame.CurrentPlayer.Name)

	x, y, o := game.MyGame.RandomLine()
	onPlayerPlayLine(x, y, o, game.MyGame.CurrentPlayer.NumPlayer)
}

func onEndGame() {
	controllerLogger.Println("Fin de la partie")

	MyTimer.Cancel()

	// Envoi de la structure de Classement
	classement := game.MyGame.GetLadder()
	SendEndGame(classement)

	// Enregistrement des scores
	game.MyGame.UpdateLadder(classement[0].Name)

	//Nouvelle instance de game
	game.StartNewGame()
}
