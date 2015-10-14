package game

import (
	"go-bourbaki/server/globals"
)

type classement []globals.PlayerClassement

// ByScore Structure héritant du tableau de joueur pour implémenter les fonctions de tri
type ByScore struct{ classement }

func (s classement) Len() int {
	return len(s)
}

func (s classement) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less Fonction de comparaison pour le tri
func (s ByScore) Less(i, j int) bool {
	return s.classement[i].Score > s.classement[j].Score
}
