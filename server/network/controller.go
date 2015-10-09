package network

import (
	"go-bourbaki/server/globals"
	"go-bourbaki/server/game"
)

func PlayLine(x int, y int, o int, n int){
	if game.MyGame.IsPauseNecessary(){
		return
	}
	l := globals.Line{x, y, o, n}
	networkLogger.Println("Un client joue en ", l)
	game.MyGame.AddLine(l)
	DisplayLine(x, y, o, n)
	isSquare, square := game.MyGame.TestSquare(l)
	if isSquare {
		networkLogger.Println("Carré")
		DisplaySquare(square.X, square.Y, square.N)
		if game.MyGame.IsEndGame() {
			//TODO appeller gestionFinPartie
		}

		// Gestion des scores
		networkLogger.Println("Attribution points")
		lastPlayer, _ := game.MyGame.GetPreviousPlayer()
		lastPlayer.Score = lastPlayer.Score - 1
		currentPlayer, _ := game.MyGame.GetPlayerFromNumPlayer(game.MyGame.CurrentPlayer.NumPlayer)
		currentPlayer.Score = currentPlayer.Score + 1
		UpdatePlayers(game.MyGame.GetAllPlayers())

		if (!game.MyGame.CurrentPlayer.IsActive){
			AI()
		}

	} else {
		game.MyGame.ChangeCurrentPlayer()
		if (!game.MyGame.CurrentPlayer.IsActive){
			AI()
		}
	}
	SetActivePlayers(game.MyGame.CurrentPlayer.NumPlayer)
}

func AI(){
	if game.MyGame.IsPauseNecessary(){
			x, y, o := game.MyGame.RandomLine()
			PlayLine(x, y, o, game.MyGame.CurrentPlayer.NumPlayer)
		}
}
