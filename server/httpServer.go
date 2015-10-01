package main

import (
	"fmt"
	"net/http"
)

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
