package network

import (
	"go-bourbaki/server/globals"

	"github.com/googollee/go-socket.io"
)

/**
 *  Liste des fonctions pour envoyer des informations au client
 *  Voir fichier notes/protocol.md pour le détail des paquets et de leurs paramètres
 */

// SendConnectAccept Packet de retour acceptant ou non la demande de connexion du joueur
// param code: int, 0 si mot de passe incorrect, 1 si connexion OK, 2 si connexion OK + compte créé
// param numPlayer: int, numéro du joueur dans la partie (entre 1 en n). 0 si le code vaut 0 (connexion refusée)
func SendConnectAccept(client socketio.Socket, code int, numPlayer int) {
	sendToClient(client, "CONNECTACCEPT", code, numPlayer)
}

// SendGrid Envoi toute la grille à un joueur (lignes et carrés)
// A appeller lors de la connexion d'un joueur
func SendGrid(client socketio.Socket, lines []globals.Line, squares []globals.Square) {
	sendToClient(client, "GRID", lines, squares)
}

// SendLadder Envoi le classement général au joueur qui l'a demandé
func SendLadder(client socketio.Socket, classementG []globals.PlayerClassement) {
	sendToClient(client, "LADDER", classementG)
}

// SendDisplayLine Ajoute la ligne x,y,o,n à la grille
func SendDisplayLine(x int, y int, orientation int, numPlayer int) {
	sendToAll("DISPLAYLINE", globals.Line{x, y, orientation, numPlayer})
}

// SendDisplaySquare Ajoute le carré x,y,n à la grille
func SendDisplaySquare(x int, y int, numPlayer int) {
	sendToAll("DISPLAYSQUARE", globals.Square{x, y, numPlayer})
}

// SendUpdatePlayers Met à jour le tableau des scores de la partie
// Appelé autant de fois que necéssaire
func SendUpdatePlayers(players []globals.Player) {
	sendToAll("UPDATEPLAYERS", players)
}

// SendSetActivePlayers Définit le joueur actif (celui qui est en train de jouer)
// Envoyé a chaque changement de joueur
func SendSetActivePlayers(currentNumPlayer int) {
	sendToAll("SETACTIVEPLAYER", currentNumPlayer)
}

// SendPause Passe la partie en pause si le nombre de joueurs actifs est inférieur à 2
func SendPause() {
	sendToAll("PAUSE")
}

// SendUnpause Sort la partie de l'état de pause
func SendUnpause() {
	sendToAll("UNPAUSE")
}

//SendEndGame Envoi le signal de fin de jeu à tous les joueurs
func SendEndGame(classement []globals.PlayerClassement) {
	sendToAll("ENDGAME", classement)
}
