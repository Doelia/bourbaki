package game

import "go-bourbaki/server/globals"

func (game *Game) getNewNumPlayer() int {
	return len(game.playersList) + 1
}

func (game *Game) addPlayer(p globals.Player) {
	game.playersList[p.Name] = p
}

func constructPlayer(numPlayer int, name string) globals.Player {
	return globals.Player{numPlayer, name, 0, true}
}
