package network

import (
	"go-bourbaki/server/globals"
	"go-bourbaki/server/game"
)

func PlayLine(x int, y int, o int, n int){
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
		UpdatePlayers(game.MyGame.GetAllPlayers())
		SetActivePlayers(game.MyGame.CurrentPlayer.NumPlayer)
	} else {
		game.MyGame.ChangeCurrentPlayer()
		SetActivePlayers(game.MyGame.CurrentPlayer.NumPlayer)
		if !game.MyGame.CurrentPlayer.IsActive{
			networkLogger.Println("pas actif")
			x, y, o := game.MyGame.RandomLine()
      PlayLine(x, y, o, game.MyGame.CurrentPlayer.NumPlayer)
		}
	}
}
