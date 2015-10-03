package network

/**
 *  Liste des fonctions pour envoyer des paquets au client
 *  Voir fichier notes/protocol.md pour le détail
 */

type linePacket struct {
	X, Y int
	O    string
	N    int
}

type squarePacket struct {
	X, Y, N int
}

type player struct {
	NumPlayer int
	Name      string
	Score     int
	IsActive  bool
}

// ConnectAccept Packet de retour acceptant ou non la demande de connexion du joueur.
func ConnectAccept(code int, numPlayer int) {
	sentToAll("CONNECTACCEPT", code, numPlayer)
}

// DisplayLine Ajoute la barre x,y,o,n à la grille
func DisplayLine(x int, y int, orientation string, numPlayer int) {
	sentToAll("CONNECTACCEPT", linePacket{x, y, orientation, numPlayer})
}

// DisplaySquare Ajoute le carré à la grille
func DisplaySquare(x int, y int, numPlayer int) {
	sentToAll("DISPLAYSQUARE", squarePacket{x, y, numPlayer})
}

// UpdatePlayers Met à jour du tableau des scores de la partie.
// Appelé autant de fois que necéssaire
func UpdatePlayers(players []player) {
	sentToAll("DISPLAYSQUARE", players)
}

// SetActivePlayers Définit le joueur actif (celui qui est en train de joueur)
// Envoyé a chaque changement de joueur
func SetActivePlayers(numPlayerActual int) {
	sentToAll("SETACTIVEPLAYER", numPlayerActual)
}
