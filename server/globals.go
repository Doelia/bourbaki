package main

import (
	"log"
	"os"
)

// Main chanel pour attente des routines
var Ch chan int

// Logger d'erreur principal
var ErrLogger = log.New(os.Stderr, "Erreur: ", 0)
