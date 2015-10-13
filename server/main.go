package main

import (
	"flag"
	"fmt"
	"go-bourbaki/server/accounts"
	"go-bourbaki/server/game"
	"go-bourbaki/server/globals"
	"go-bourbaki/server/network"
	"math/rand"
	"time"
)

var test = flag.String("test", "main", "Sélectionne la méthode de test à lancer (debug uniquement)")
var port = flag.Int("port", 2000, "Modifie le port d'écoute (défaut 2000)")

func goMain() {
	rand.Seed(time.Now().Unix())
	globals.Ch = make(chan int, 1)

	fmt.Println("=== BOURBAKI SERVEUR ===")

	// Initialisation de la base de donnée
	accounts.OpenDB()

	// Création d'une partie
	game.StartNewGame()

	// Création serveur HTTP
	go network.StartWebServer(*port)
	<-globals.Ch // Attente de l'handle

	<-globals.Ch // Attente fin serveur http (ne doit pas arriver)
}

func main() {

	flag.Parse()

	switch *test {
	case "main":
		goMain()
	case "db":
		// Initialisation de la base de donnée
		accounts.OpenDB()
		accounts.Testdb()
	}

}
