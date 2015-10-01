package main

import "fmt"

func main() {

	Ch = make(chan int, 1)

	fmt.Println("=== BOURBAKI SERVEUR ===")

	// Création serveur HTTP
	go startHTTPServer(2000)
	<-Ch // Wait handle HTTP

	fmt.Println("Next")

	<-Ch // Wait HTTP Server
}
