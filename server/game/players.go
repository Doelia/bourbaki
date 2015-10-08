package game

import (
	"go-bourbaki/server/globals"
	"errors"
)

// GetNewNumPlayer Retourne un numéro de joueur pour le joueur suivant
func (g *Game) GetNewNumPlayer() int {
	return len(g.playersList) + 1
}

// AddPlayer Ajoute un joueur au game
func (g *Game) AddPlayer(p globals.Player) {
	g.playersList[p.Name] = p
	gameLogger.Println("Ajout du joueur " + p.Name + " à la partie")
}

// GetPlayer Retourne une structure du player demandé
func (g *Game) GetPlayerFromName(name string) (globals.Player, error) {
	player, exists := g.playersList[name]
	if !exists {
		return globals.Player{}, errors.New("Joueur introuvable")
	}
	return player, nil
}

func (g *Game) GetPlayerFromNumPlayer(numPlayer int)(globals.Player, error){
	for _, playerStruct := range g.playersList{
		if playerStruct.NumPlayer == numPlayer{
			return playerStruct, nil
		}
	}
	return globals.Player{}, errors.New("Joueur introuvable")
}

// PlayerExists Retourne vrai si le joueur existe dans la partie
func (g *Game) PlayerExists(name string) bool {
	_, ok := g.playersList[name]
	return ok
}

// ConstructPlayer Construit un nouveau Player correctement initialisé
func ConstructPlayer(numPlayer int, name string) globals.Player {
	return globals.Player{numPlayer, name, 0, true}
}
