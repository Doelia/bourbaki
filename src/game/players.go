package game

import (
	"errors"
	"github.com/doelia/go-bourbaki/src/globals"
)

// GetNewNumPlayer Retourne un numéro de joueur pour le joueur suivant
func (g *Game) GetNewNumPlayer() int {
	return len(g.playersList) + 1
}

// GetPreviousPlayer permet de récupérer le joueur précédent
// @return un pointeur vers le joueur
func (g *Game) GetPreviousPlayer() (*globals.Player, error) {
	if g.CurrentPlayer.NumPlayer == 1 {
		return g.GetPlayerFromNumPlayer(len(g.playersList))
	}
	return g.GetPlayerFromNumPlayer(g.CurrentPlayer.NumPlayer - 1)
}

// AddPlayer Ajoute un joueur au game
func (g *Game) AddPlayer(p globals.Player) {
	g.playersList[p.Name] = &p
	gameLogger.Println("Ajout du joueur " + p.Name + " à la partie")
}

// GetAllPlayers Retoure un tableau des joueurs de la partie triés par leur numéro de joueur
// Utile pour envoyer un json
func (g *Game) GetAllPlayers() []globals.Player {
	var list []globals.Player
	for _, p := range g.playersList {
		list = append(list, *p)
	}
	return list
}

// GetPlayerFromName Retourne une structure du player demandé
func (g *Game) GetPlayerFromName(name string) (*globals.Player, error) {
	_, exists := g.playersList[name]
	if !exists {
		return nil, errors.New("Joueur introuvable")
	}
	return g.playersList[name], nil
}

// GetPlayerFromNumPlayer Retourne un joueur à partir de son numéro dans la partie
func (g *Game) GetPlayerFromNumPlayer(numPlayer int) (*globals.Player, error) {
	for _, playerStruct := range g.playersList {
		if playerStruct.NumPlayer == numPlayer {
			return playerStruct, nil
		}
	}
	return nil, errors.New("Joueur introuvable")
}

// GetPlayerFromIDSocket Retourne un joueur à partir de son id de socket
func (g *Game) GetPlayerFromIDSocket(idSocket string) (*globals.Player, error) {
	for _, playerStruct := range g.playersList {
		if playerStruct.IDSocket == idSocket {
			return playerStruct, nil
		}
	}
	return nil, errors.New("Joueur introuvable")
}

// PlayerExists Retourne vrai si le joueur existe dans la partie
func (g *Game) PlayerExists(name string) bool {
	_, ok := g.playersList[name]
	return ok
}

// ConstructPlayer Construit un nouveau Player correctement initialisé
func ConstructPlayer(numPlayer int, name string, idSocket string) globals.Player {
	return globals.Player{numPlayer, name, 0, true, idSocket}
}
