package main

import (
	"fmt"
	"net/http"

	"github.com/googollee/go-socket.io"
)

func startWebServer(port int, io *socketio.Server) {
	fmt.Printf("Démarrage serveur web sur le port %d...\n", port)
	http.Handle("/", http.FileServer(http.Dir("../client")))
	http.Handle("/socket.io/", io)
	Ch <- 1
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		ErrLogger.Println("Erreur à la création du serveur HTTP : ", err.Error())
	}
	Ch <- 1
}
