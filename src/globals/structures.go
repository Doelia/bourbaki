package globals

// Line Représente une ligne sur le plateau (active si n!=0)
type Line struct {
	X, Y, O, N int
}

// Square Représente un carré sur le plateu (actif si n!=0)
type Square struct {
	X, Y, N int
}

// HORIZONTAL code int pour les lignes horizontales
const HORIZONTAL = 0

// VERTICAL code int pour les lignes verticales
const VERTICAL = 1

// Player structure définissant un joueur
type Player struct {
	NumPlayer int
	Name      string
	Score     int
	IsActive  bool
	IDSocket  string
}

// PlayerClassement structure définissant un joueur dans le tableau du classement
type PlayerClassement struct {
	Classement int
	NumPlayer  int
	Name       string
	Score      int
	NbrGames   int
	NbrWins    int
}
