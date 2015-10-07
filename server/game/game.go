package game

import (
	"go-bourbaki/server/globals"
)

// Game structure définissant une partie
type Game struct {
	lines       [globals.GRIDSIZE][globals.GRIDSIZE][2]int
	squares     [globals.GRIDSIZE][globals.GRIDSIZE]int
	playersList map[string]globals.Player
}

// ConstructGame Construit et initialise un nouveau jeu
func ConstructGame() *Game {
	var game = &Game{}
	game.playersList = make(map[string]globals.Player)
	return game
}

var myGame *Game

// TestGame ..
func TestGame() {
	myGame = ConstructGame()
	player := constructPlayer(myGame.getNewNumPlayer(), "Pancake")
	myGame.addPlayer(player)
}
