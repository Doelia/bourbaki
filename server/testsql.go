package main

import (
	"fmt"
	"github.com/boltdb/bolt"
)

func Testsql() {
	db, err := bolt.Open("bourbaki.db", 0600, nil)
	if err != nil {
		fmt.Print("Erreur")
	}
	defer db.Close()
}
