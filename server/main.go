package main

import (
	"flag"
	"fmt"
	"go-bourbaki/server/accounts"
	"go-bourbaki/server/game"
	"go-bourbaki/server/globals"
	"go-bourbaki/server/network"
)

var test = flag.String("test", "main", "Sélectionne la méthode de test à lancer (debug uniquement)")
var port = flag.Int("port", 2000, "Modifie le port d'écoute (défaut 2000)")

func dbTest() {
	accounts.Testsql()
}

func serverTest() {
	globals.Ch = make(chan int, 1)

	fmt.Println("=== BOURBAKI SERVEUR ===")

	// Création de gane
	game.StartNewGame()

	// Création serveur HTTP
	go network.StartWebServer(*port)
	<-globals.Ch // Attente de l'handle

	<-globals.Ch // Attente fin serveur http (ne doit pas arriver)
}

func main() {
	accounts.OpenDB()
	flag.Parse()

	switch *test {
	case "db":
		dbTest()
	case "game":
		game.TestGame()
	case "main":
		serverTest()
	}

}
