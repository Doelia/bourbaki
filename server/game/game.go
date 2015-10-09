package game

import (
	"go-bourbaki/server/globals"
	"log"
	"os"
)

var gameLogger = log.New(os.Stdout, "[game] ", 0)

// Game structure définissant une partie
type Game struct {
	lines         [globals.GRIDSIZE][globals.GRIDSIZE][2]int
	squares       [globals.GRIDSIZE][globals.GRIDSIZE]int
	playersList   map[string]*globals.Player
	CurrentPlayer globals.Player
}

// MyGame variable globable de l'instance unique d'une partie
var MyGame *Game

// ConstructGame Construit et initialise un nouveau jeu
func ConstructGame() *Game {
	var game = &Game{}
	game.playersList = make(map[string]*globals.Player)
	return game
}

// StartNewGame Démarre une nouvelle partie en initialisant tous les structures associées
func StartNewGame() {
	gameLogger.Println("Création d'une nouvelle partie...")
	MyGame = ConstructGame()
	gameLogger.Println("Création OK")
}

// AddLine Active la ligne dans le game
func (g *Game) AddLine(line globals.Line) {
	if g.lines[line.X][line.Y][line.O] == 0 {
		g.lines[line.X][line.Y][line.O] = line.N
	}
}

// AddSquare Active le carré dans le game
func (g *Game) AddSquare(square globals.Square) {
	g.squares[square.X][square.Y] = square.N
}

// isActive Retourne vrai si la ligne est active dans le game, faux sinon
func (g *Game) isActive(x int, y int, o int) bool {
	if x < 0 || x >= globals.GRIDSIZE {
		return false
	}
	if y < 0 || y >= globals.GRIDSIZE {
		return false
	}
	return g.lines[x][y][o] > 0
}

// TestSquare permet de savoir si la ligne qui vient d'être jouée forme un carré
// @param lastLine: dernière ligne ayant été jouée
// @return bool: vrai si le joueur gagne un carré, faux sinon
// @return square: le carré formé
func (g *Game) TestSquare(lastLine globals.Line) (bool, globals.Square) {
	x := lastLine.X
	y := lastLine.Y
	if lastLine.O == globals.HORIZONTAL {
		if g.isActive(x, y-1, globals.HORIZONTAL) && g.isActive(x+1, y-1, globals.VERTICAL) && g.isActive(x, y-1, globals.VERTICAL) {
			gameLogger.Println("Ajout square au dessus du trait")
			return true, globals.Square{x, y - 1, lastLine.N}
		}
		if g.isActive(x, y+1, globals.HORIZONTAL) && g.isActive(x, y, globals.VERTICAL) && g.isActive(x+1, y, globals.VERTICAL) {
			gameLogger.Println("Ajout square au dessous du trait")
			return true, globals.Square{x, y, lastLine.N}
		}
	} else {
		if lastLine.O == globals.VERTICAL {
			if g.isActive(x, y, globals.HORIZONTAL) && g.isActive(x+1, y, globals.VERTICAL) && g.isActive(x, y+1, globals.HORIZONTAL) {
				gameLogger.Println("Ajout square à droite du trait")
				return true, globals.Square{x, y, lastLine.N}
			}
			if g.isActive(x-1, y, 0) && g.isActive(x-1, y, 1) && g.isActive(x-1, y+1, 0) {
				gameLogger.Println("Ajout square à gauche du trait")
				return true, globals.Square{x - 1, y, lastLine.N}
			}
		}
	}
	gameLogger.Println("Pas de square")
	return false, globals.Square{}
}

// ChangeCurrentPlayer permet de changer le joueur courant, à appeller lors de la fin d'un tour
func (g *Game) ChangeCurrentPlayer() {
	numNewCurrentPlayer := g.CurrentPlayer.NumPlayer + 1
	if numNewCurrentPlayer > len(g.playersList) {
		numNewCurrentPlayer = 1
	}
	newCurrentPlayer, err := g.GetPlayerFromNumPlayer(numNewCurrentPlayer)
	if err != nil {
		gameLogger.Println("Changement joueur courant impossible")
	}
	if newCurrentPlayer.IsActive {
		g.CurrentPlayer = *newCurrentPlayer
		gameLogger.Println("Joueur courant : ", g.CurrentPlayer.Name)
	} else {
		g.ChangeCurrentPlayer()
	}
}

// IsPauseNecessary permet de savoir si une pause est nécessaire (nbJoueursActifs >= 2)
func (g *Game) IsPauseNecessary() bool {
	//on compte le nombre de joueurs actifs
	compteur := 0
	for _, playerStruct := range g.playersList {
		if playerStruct.IsActive == true {
			compteur++
		}
	}
	return compteur < 2
}

// IsEndGame permet de savoir si la partie est finie
// TODO EndGameManager
func (g *Game) IsEndGame() bool {
	for i := 0; i < len(g.squares)-1; i++ {
		for j := 0; j < len(g.squares)-1; j++ {
			if g.squares[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
