package main

import (
	"encoding"
	"fmt"
	"github.com/boltdb/bolt"
)

var db *bolt.DB     //bd
var tx *bolt.Tx     //transaction
var b *bolt.Bucket  //bucket
var br *bolt.Bucket //bucket
var erreur error

func OpenDB() {
	db, err := bolt.Open("bourbaki.db", 0600, nil)
	if err != nil {
		fmt.Print("Erreur à l'ouverture de la base de données")
	}
	defer db.Close()
}

func OpenTransaction() {
	tx, err2 := db.Begin(true)
	if err2 != nil {
		fmt.Println("Lancement de la transaction")
	}
	defer tx.Rollback()
}

func CreateBucket() {
	b, err3 := tx.CreateBucketIfNotExists([]byte("Accounts"))
	if err3 != nil {
		fmt.Println("Erreur création bucket Accounts")
	}
}

func OpenBucket() {
	br := tx.Bucket([]byte("MyBucket"))
}

// Ajoute a dans la db. Retourne vrai si ça a fonctionné, faux sinon
func AddInDB(a Account) bool {
	//il ne reste plus qu'à transformer le account en chaîne de carac
	err4 := b.Put([]byte(a.Name), []byte(a))
	if err4 != nil {
		fmt.Println("Erreur put")
		return false
	}
	return true
}

//Affiche
func getFromDB(cle string) Account {
	v := br.Get([]byte(cle))
	return v
}

func Testsql() {
	OpenDB()
	OpenTransaction()
	CreateBucket()
	//Il faut appeller AddInDB avec un account
}
