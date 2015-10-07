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

func (g *Game) addLine(line globals.Line){
	g.lines[line.X][line.Y][line.O] = line.N
}

func (g *Game) addSquare(square globals.Square){
	g.squares[square.X][square.Y] = square.N
}

// testSquare permet de savoir si la ligne qui vient d'être jouée forme un carré
//@param lastLine: dernière ligne ayant été jouée
//@return bool: vrai si le joueur gagne un carré, faux sinon
//@return square: le carré formé
func (g *Game) testSquare(lastLine globals.Line) (bool, globals.Square){
	x := lastLine.X
	y := lastLine.Y
	if lastLine.O == 0{	// ligne horizontale
		if g.isActive(x, y-1, 0) && g.isActive(x+1, y-1, 1) && g.isActive(x, y-1, 1){
			return true, globals.Square{x, y-1, lastLine.N}
		}
		if g.isActive(x, y+1, 0) && g.isActive(x, y, 1) && g.isActive(x+1, y, 1){
			return true, globals.Square{x, y, lastLine.N}
		}
	} else {	// ligne verticale
		if g.isActive(x, y, 0) && g.isActive(x+1, y, 1) && g.isActive(x, y+1, 0){
			return true, globals.Square{x, y, lastLine.N}
		}
		if g.isActive(x-1, y, 0) && g.isActive(x-1, y, 1) && g.isActive(x-1, y+1, 0){
			return true, globals.Square{x-1, y, lastLine.N}
		}
	}
	return false, globals.Square{}
}

// fonction qui retourne vrai si la ligne est active dans g, faux sinon
func (g *Game) isActive(x int, y int, o int) bool{
	if g.lines[x][y][o] == 0{
		return false
	}
	return true
}
