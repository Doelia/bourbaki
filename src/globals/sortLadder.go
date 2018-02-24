package globals

// Classement Alias pour un tableau de PlayerClassement
type Classement []PlayerClassement

// ByScore Structure héritant du tableau de joueur pour implémenter les fonctions de tri
type ByScore struct{ Classement }

// Len Fonction de surcharge pour le sort
func (s Classement) Len() int {
	return len(s)
}

// Swap Fonction de surcharge pour le sort
func (s Classement) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less Fonction de comparaison pour le tri
func (s ByScore) Less(i, j int) bool {
	return s.Classement[i].Score > s.Classement[j].Score
}
