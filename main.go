package main

import (
	"flag"
	"fmt"
	"github.com/doelia/go-bourbaki/src/accounts"
	"github.com/doelia/go-bourbaki/src/game"
	"github.com/doelia/go-bourbaki/src/globals"
	"github.com/doelia/go-bourbaki/src/network"
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
	network.OnCreateGame()

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
	}

}
