package globals

// Line ..
type Line struct {
	X, Y int
	O    int
	N    int
}

// Square ..
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
