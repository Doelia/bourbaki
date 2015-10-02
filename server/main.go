package main

import (
	"fmt"
	"go-bourbaki/server/accounts"
)

func main() {

	accounts.Testsql()

	Ch = make(chan int, 1)

	fmt.Println("=== BOURBAKI SERVEUR ===")

	// Création serveur HTTP
	go startWebServer(2000, getWebSocketHandler())
	<-Ch // Wait handle HTTP

	fmt.Println("Next")

	<-Ch // Wait HTTP Server
}
