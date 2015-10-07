package game

import (
	"go-bourbaki/server/globals"
	"log"
	"os"
)

var gameLogger = log.New(os.Stdout, "[game] ", 0)

// Game structure définissant une partie
type Game struct {
	lines       [globals.GRIDSIZE][globals.GRIDSIZE][2]int
	squares     [globals.GRIDSIZE][globals.GRIDSIZE]int
	playersList map[string]globals.Player
}

// MyGame variable globable de l'instance unique d'une partie
var MyGame *Game

// ConstructGame Construit et initialise un nouveau jeu
func ConstructGame() *Game {
	var game = &Game{}
	game.playersList = make(map[string]globals.Player)
	return game
}

// StartNewGame Démarre une nouvelle partie en initialisant tous les structure associés
func StartNewGame() {
	gameLogger.Println("Création d'une nouvelle partie...")
	MyGame = ConstructGame()
	gameLogger.Println("Création OK")
}

// TestGame ..
func TestGame() {
	player := ConstructPlayer(MyGame.GetNewNumPlayer(), "Pancake")
	MyGame.AddPlayer(player)
}
