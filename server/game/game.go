package game

import (
	"go-bourbaki/server/globals"
	"go-bourbaki/server/accounts"
	"log"
	"math/rand"
	"os"
	"sort"
)

var gameLogger = log.New(os.Stdout, "[game] ", 0)

// Game classe définissant une partie
type Game struct {
	lines         [globals.GRIDSIZE][globals.GRIDSIZE][2]int
	squares       [globals.GRIDSIZE][globals.GRIDSIZE]int
	playersList   map[string]*globals.Player
	CurrentPlayer *globals.Player
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
	gameLogger.Println("Création OK. En attente de joueurs.")
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
	g.CurrentPlayer = newCurrentPlayer
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
	if g.CurrentPlayer.NumPlayer == 1 {
		return g.GetPlayerFromNumPlayer(len(g.playersList))
	}
	return g.GetPlayerFromNumPlayer(g.CurrentPlayer.NumPlayer - 1)
}

// RandomLine Retoune une ligne aléatoire du plateau pas encore jouée
func (g *Game) RandomLine() (int, int, int) {
	for {
		i := rand.Intn(globals.GRIDSIZE)
		j := rand.Intn(globals.GRIDSIZE)
		k := rand.Intn(2)
		if g.lines[i][j][k] == 0 {
			return i, j, k
		}
	}
}

// SaveScores TODO commentaire
func (g *Game) UpdateLadder(nameGagnant string) {
	for _, player := range g.playersList {
		account := accounts.GetFromDB(player.Name)
		account.Points += player.Score
		account.NbrGames++
		if (account.Name == nameGagnant){
			account.NbrWins++
		}
		accounts.UpdateAccount(account)
		gameLogger.Println(player.Name, " gagne ", account.Points, " points")
	}
}

// GetLadder TODO commentaire
func (g *Game) GetLadder() globals.Classement {
	var classementtb globals.Classement
	// 1e étape: récupération du classement
	for _, player := range g.playersList {
		p := globals.PlayerClassement{0, player.NumPlayer, player.Name, player.Score, 0, 0}
		classementtb = append(classementtb, p)
	}

	// 2e étape: tri par Score
	sort.Sort(globals.ByScore{classementtb})

	// 3e étape: ajout de l'attribut Classement
	for i := 1; i <= len(classementtb); i++ {
		classementtb[i-1].Classement = i
	}
	gameLogger.Println("Classement: ", classementtb)

	return classementtb
}
