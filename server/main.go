package main

import (
	"flag"
	"fmt"
	"go-bourbaki/server/accounts"
	"go-bourbaki/server/globals"
	"go-bourbaki/server/network"
	"time"
)

var test = flag.String("test", "main", "Selectionne la méthode de test à lancer")
var port = flag.Int("port", 2000, "Modifie le port d'écoute (défaut 2000)")

func dbTest() {
	accounts.Testsql()
}

func serverTest() {
	globals.Ch = make(chan int, 1)

	fmt.Println("=== BOURBAKI SERVEUR ===")

	// Création serveur HTTP
	go network.StartWebServer(*port)
	<-globals.Ch // Attente de l'handle

	for {
		network.DisplayLine(3, 3, "v", 1)
		time.Sleep(3 * time.Second)
	}

	<-globals.Ch // Attente fin serveur http (ne doit pas arriver)
}

func main() {

	flag.Parse()

	switch *test {
	case "db":
		dbTest()
	case "main":
		serverTest()
	}

}
