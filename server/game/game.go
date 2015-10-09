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

// GetPreviousPlayer permet de récupérer le joueur précédent
func (g *Game) GetPreviousPlayer() (*globals.Player, error) {
	if g.CurrentPlayer.NumPlayer == len(g.playersList) {
		return g.GetPlayerFromNumPlayer(1)
	}
	return g.GetPlayerFromNumPlayer(g.CurrentPlayer.NumPlayer + 1)
}

// TODO recherche random
func (g *Game) RandomLine() (int, int, int){
	for i := 0; i < len(g.lines); i++{
		for j := 0; j < len(g.lines); j++{
			for k := 0; k < 2; k++{
				if g.lines[i][j][k] == 0{
					return i,j,k
				}
			}
		}
	}
}
