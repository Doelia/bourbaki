package network

import (
	"go-bourbaki/server/game"
	"go-bourbaki/server/globals"
	"log"
	"os"

	"github.com/googollee/go-socket.io"
)

var controllerLogger = log.New(os.Stdout, "[event] ", 0)

func onResume() {
	controllerLogger.Println("onResume()")
	Unpause()
	onNewTurn()
}

func onPause() {
	controllerLogger.Println("onPause()")
	Pause()
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
	}
	if game.MyGame.CurrentPlayer == nil {
		game.MyGame.CurrentPlayer = player
	}
	player.IsActive = true
	ConnectAccept(so, resultatIntLogin, numPlayer)

	// Suivera d'un onReady()
}

func onReady(so socketio.Socket) {
	controllerLogger.Println("onReady")

	Grid(so, game.MyGame.GetActivesLinesList(), game.MyGame.GetActivesSquaresList())
	UpdatePlayers(game.MyGame.GetAllPlayers())
	SetActivePlayers(game.MyGame.CurrentPlayer.NumPlayer)

	if !game.MyGame.IsPauseNecessary() {
		onResume()
	} else {
		onPause()
	}
}

func onLeft(player *globals.Player) {
	controllerLogger.Println("onLeft: ", player.Name)

	player.IsActive = false
	UpdatePlayers(game.MyGame.GetAllPlayers())

	if game.MyGame.IsPauseNecessary() {
		onPause()
	} else {
		if player.NumPlayer == game.MyGame.CurrentPlayer.NumPlayer {
			AI()
		}
	}

}

func onSquareDone(square globals.Square) {
	controllerLogger.Println("onSquareDone: ", square)

	DisplaySquare(square.X, square.Y, square.N)

	// Gestion des scores
	lastPlayer, _ := game.MyGame.GetPreviousPlayer()
	lastPlayer.Score = lastPlayer.Score - 1
	currentPlayer, _ := game.MyGame.GetPlayerFromNumPlayer(game.MyGame.CurrentPlayer.NumPlayer)
	currentPlayer.Score = currentPlayer.Score + 1

	UpdatePlayers(game.MyGame.GetAllPlayers())
}

func onNewTurn() {
	controllerLogger.Println("onNewTurn: ", game.MyGame.CurrentPlayer)

	controllerLogger.Println(game.MyGame.GetAllPlayers())

	if !game.MyGame.CurrentPlayer.IsActive {
		AI()
	} else {
		SetActivePlayers(game.MyGame.CurrentPlayer.NumPlayer)
	}
}

func onPlayerPlayLine(x int, y int, o int, n int) {

	if game.MyGame.IsPauseNecessary() {
		return
	}

	l := globals.Line{x, y, o, n}
	controllerLogger.Println("onPlayerPlayLine: ", game.MyGame.CurrentPlayer.Name, l)

	game.MyGame.AddLine(l)
	DisplayLine(x, y, o, n)

	isSquare, square := game.MyGame.TestSquare(l)
	if isSquare {
		onSquareDone(square)
	} else {
		game.MyGame.ChangeCurrentPlayer()
	}

	onNewTurn()

}

// AI ...
func AI() {
	controllerLogger.Println("AI pour " + game.MyGame.CurrentPlayer.Name)

	x, y, o := game.MyGame.RandomLine()
	onPlayerPlayLine(x, y, o, game.MyGame.CurrentPlayer.NumPlayer)
}
