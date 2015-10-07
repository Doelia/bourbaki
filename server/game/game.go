package game

import (
	"go-bourbaki/server/globals"
)

// Game structure définissant une partie
type Game struct {
	lines   [globals.GRIDSIZE][globals.GRIDSIZE][2]int
	squares [globals.GRIDSIZE][globals.GRIDSIZE]int
}

// Line ..
type Line struct {
	X, Y int
	O    string
	N    int
}

// Square ..
type Square struct {
	X, Y, N int
}

// Player structure définissant un joueur
type Player struct {
	NumPlayer int
	Name      string
	Score     int
	IsActive  bool
}
