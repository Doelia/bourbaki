package network

import (
	"fmt"
	"go-bourbaki/server/globals"
	"log"
	"net/http"
	"os"
)

var networkLogger = log.New(os.Stdout, "[network] ", 0)

// StartWebServer démarre le serveur web (http + websockets)
func StartWebServer(port int) {
	networkLogger.Printf("Démarrage du serveur web sur le port %d...\n", port)
	http.Handle("/", http.FileServer(http.Dir("../client")))
	http.Handle("/socket.io/", createWebSocketHandler())
	globals.Ch <- 1
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		globals.ErrLogger.Println("Erreur à la création du serveur HTTP : ", err.Error())
	}
	globals.Ch <- 1
}
