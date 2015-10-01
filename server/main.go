package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Main chanel
var Ch chan int

// Logger d'erreur principal
var ErrLogger = log.New(os.Stderr, "Erreur: ", 0)

func startHTTPServer(port int) {
	fmt.Printf("Start http server on port %d...\n", port)
	http.Handle("/", http.FileServer(http.Dir("../client")))
	Ch <- 1
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		ErrLogger.Println("Erreur à la création du serveur HTTP : ", err.Error())
	}
	Ch <- 1
}

func main() {

	Ch = make(chan int, 1)

	fmt.Println("=== BOURBAKI SERVEUR ===")

	// Création serveur HTTP
	go startHTTPServer(2000)
	<-Ch // Wait handle HTTP

	fmt.Println("Next")

	<-Ch // Wait HTTP Server
}
