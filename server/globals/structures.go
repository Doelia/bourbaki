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

// Player structure définissant un joueur
type Player struct {
	NumPlayer int
	Name      string
	Score     int
	IsActive  bool
}
