package game

import (
	"go-bourbaki/server/globals"
  "go-bourbaki/server/network"
)

func PlayLine(x int, y int, o int, n int){
	l := globals.Line{x, y, o, n}
	gameLogger.Println("Un client joue en ", l)
	MyGame.AddLine(l)
	network.DisplayLine(x, y, o, n)
	isSquare, square := MyGame.TestSquare(l)
	if isSquare {
		network.DisplaySquare(square.X, square.Y, square.N)
		if MyGame.IsEndGame() {
			//TODO appeller gestionFinPartie
		}
		lastPlayer, _ := MyGame.GetPreviousPlayer()
		lastPlayer.Score = lastPlayer.Score - 1
		currentPlayer, _ := MyGame.GetPlayerFromNumPlayer(MyGame.CurrentPlayer.NumPlayer)
		currentPlayer.Score = currentPlayer.Score + 1
		network.UpdatePlayers(MyGame.GetAllPlayers())
		network.SetActivePlayers(MyGame.CurrentPlayer.NumPlayer)
	} else {
		MyGame.ChangeCurrentPlayer()
		network.SetActivePlayers(MyGame.CurrentPlayer.NumPlayer)
		if !MyGame.CurrentPlayer.IsActive{
			x, y, o := MyGame.RandomLine()
      PlayLine(x, y, o , MyGame.CurrentPlayer.NumPlayer)
		}
	}
}
