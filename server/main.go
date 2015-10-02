package main

import (
	"fmt"
	"go-bourbaki/server/accounts"
	"flag"
	"go-bourbaki/server/globals"
	"go-bourbaki/server/network"
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
	<-globals.Ch // Wait handle HTTP

	<-globals.Ch // Wait HTTP Server
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
