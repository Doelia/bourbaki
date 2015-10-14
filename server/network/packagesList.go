package network

import (
	"go-bourbaki/server/globals"

	"github.com/googollee/go-socket.io"
)

/**
 *  Liste des fonctions pour envoyer des informations au client
 *  Voir fichier notes/protocol.md pour le détail des paquets et de leurs paramètres
 */

// SendConnectAccept Packet de retour acceptant ou non la demande de connexion du joueur.
// param code: int, 0 si mot de passe incorrect, 1 si connexion OK, 2 si connexion OK + compte créé
// param numPlayer: int, numéro du joueur dans la partie (entre 1 en n). 0 si le code vaut 0 (connexion refusée)
func SendConnectAccept(client socketio.Socket, code int, numPlayer int) {
	sendToClient(client, "CONNECTACCEPT", code, numPlayer)
}

// SendGrid Envoi toute la grille à un joueur (lignes et carrés)
// A envoyer quand il se connecte
func SendGrid(client socketio.Socket, lines []globals.Line, squares []globals.Square) {
	sendToClient(client, "GRID", lines, squares)
}

// SendDisplayLine Ajoute la barre x,y,o,n à la grille
func SendDisplayLine(x int, y int, orientation int, numPlayer int) {
	sentToAll("DISPLAYLINE", globals.Line{x, y, orientation, numPlayer})
}

// SendDisplaySquare Ajoute le carré à la grille
func SendDisplaySquare(x int, y int, numPlayer int) {
	sentToAll("DISPLAYSQUARE", globals.Square{x, y, numPlayer})
}

// SendUpdatePlayers Met à jour du tableau des scores de la partie
// Appelé autant de fois que necéssaire
func SendUpdatePlayers(players []globals.Player) {
	sentToAll("UPDATEPLAYERS", players)
}

// SendSetActivePlayers Définit le joueur actif (celui qui est en train de joueur)
// Envoyé a chaque changement de joueur
func SendSetActivePlayers(currentNumPlayer int) {
	sentToAll("SETACTIVEPLAYER", currentNumPlayer)
}

// SendPause Passe la partie en pause si le nombre de joueurs actifs est inférieur à 2
func SendPause() {
	sentToAll("PAUSE")
}

// SendUnpause Sort la partie de l'état de pause
func SendUnpause() {
	sentToAll("UNPAUSE")
}


func SendEndGame(classement []globals.PlayerClassement){
	sentToAll("ENDGAME", classement)
}
