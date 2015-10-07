package game

import (
	"go-bourbaki/server/globals"
)

// Line structure
type Line struct {
	X, Y int // coordonnées (x,y) de la ligne
	O    int // orientation de la ligne
	N    int // num joueur
}

// Square structure
type Square struct {
	X, Y int // (x,y) coordonnées du point origine du carré
	N int // num joueur
}

// Player structure définissant un joueur
type Player struct {
	NumPlayer int
	Name      string
	Score     int
	IsActive  bool
}

// Game classe définissant une partie
type Game struct {
	lines   [globals.GRIDSIZE][globals.GRIDSIZE][2]int
	squares [globals.GRIDSIZE][globals.GRIDSIZE]int
}

func (g *Game) addLine(line Line){
	g.lines[line.X][line.Y][line.O] = line.N
}

func (g *Game) addSquare(square Square){
	g.squares[square.X][square.Y] = square.N
}

// testSquare permet de savoir si la ligne qui vient d'être jouée forme un carré
//@param lastLine: dernière ligne ayant été jouée
//@return bool: vrai si le joueur gagne un carré, faux sinon
//@return square: le carré formé
func (g *Game) testSquare(lastLine Line) (bool, Square){
	x := lastLine.X
	y := lastLine.Y
	if lastLine.O == 0{	// ligne horizontale
		if g.isActive(x, y-1, 0) && g.isActive(x+1, y-1, 1) && g.isActive(x, y-1, 1){
			return true, Square{x, y-1, lastLine.N}
		}
		if g.isActive(x, y+1, 0) && g.isActive(x, y, 1) && g.isActive(x+1, y, 1){
			return true, Square{x, y, lastLine.N}
		}
	} else {	// ligne verticale
		if g.isActive(x, y, 0) && g.isActive(x+1, y, 1) && g.isActive(x, y+1, 0){
			return true, Square{x, y, lastLine.N}
		}
		if g.isActive(x-1, y, 0) && g.isActive(x-1, y, 1) && g.isActive(x-1, y+1, 0){
			return true, Square{x-1, y, lastLine.N}
		}
	}
	return false, Square{}
}

// fonction qui retourne vrai si la ligne est active dans g, faux sinon
func (g *Game) isActive(x int, y int, o int) bool{
	if g.lines[x][y][o] == 0{
		return false
	}
	return true
}
