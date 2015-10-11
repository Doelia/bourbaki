package globals

import (
	"log"
	"os"
)

// Main chanel pour attente des routines
var Ch chan int

// Logger d'erreur principal
var ErrLogger = log.New(os.Stderr, "[error] ", 0)

// Taille de la grille de jeu (nombre de points)
// Note: Modifier aussi dans les fichiers JS du client
const GRIDSIZE = 10
