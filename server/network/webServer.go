package network

import (
	"fmt"
	"go-bourbaki/server/globals"
	"net/http"
)

// StartWebServer démarre le serveur web (http + websockets)
func StartWebServer(port int) {
	fmt.Printf("Démarrage serveur web sur le port %d...\n", port)
	http.Handle("/", http.FileServer(http.Dir("../client")))
	http.Handle("/socket.io/", getWebSocketHandler())
	globals.Ch <- 1

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		globals.ErrLogger.Println("Erreur à la création du serveur HTTP : ", err.Error())
	}
	globals.Ch <- 1
}
