package game

import "go-bourbaki/server/globals"

// GetNewNumPlayer ..
func (game *Game) GetNewNumPlayer() int {
	return len(game.playersList) + 1
}

// AddPlayer ..
func (game *Game) AddPlayer(p globals.Player) {
	game.playersList[p.Name] = p
	gameLogger.Println("Ajout du joueur " + p.Name + " à la partie")
}

// PlayerExists ..
func (game *Game) PlayerExists(name string) bool {
	_, ok := game.playersList[name]
	return ok
}

// ConstructPlayer ..
func ConstructPlayer(numPlayer int, name string) globals.Player {
	return globals.Player{numPlayer, name, 0, true}
}
