package main

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
)

var db *bolt.DB    //bd
var tx *bolt.Tx    //transaction
var b *bolt.Bucket //bucket
var err error

func OpenDB() {
	db, err = bolt.Open("bourbaki.db", 0600, nil)
	if err != nil {
		fmt.Print("Erreur à l'ouverture de la base de données")
	}
}

func OpenTransaction() {
	tx, err = db.Begin(true)
	if err != nil {
		fmt.Println("Lancement de la transaction")
	}
}

func EndTransaction() {
	tx.Commit()
}

func CreateBucket() {
	OpenTransaction()
	defer EndTransaction()
	b, err = tx.CreateBucketIfNotExists([]byte("Accounts"))
	if err != nil {
		fmt.Println("Erreur création bucket Accounts")
	}
}

// Ajoute a dans la db. Retourne vrai si ça a fonctionné, faux sinon
func AddInDB(a Account) bool {
	OpenTransaction()
	defer EndTransaction()
	astring, _ := json.Marshal(a)

	err = b.Put([]byte(a.Name), []byte(astring))
	if err != nil {
		fmt.Println("Erreur put")
		return false
	}
	return true
}

// Affiche
func getFromDB(cle string) Account {
	OpenTransaction()
	defer EndTransaction()
	v := b.Get([]byte(cle))
	var res Account
	json.Unmarshal(v, &res)
	return res
}

func Testsql() {
	OpenDB()
	CreateBucket()
	a := Account{"name: naaaaame", "pass:paaaass", 12}
	AddInDB(a)
	//Il faut appeller AddInDB avec un account
}
