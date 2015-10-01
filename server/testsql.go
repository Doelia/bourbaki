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

	tx, err2 := db.Begin(true)
	if err2 != nil {
		fmt.Println("Erreur création db")
	}
	defer tx.Rollback()

	b, err3 := tx.CreateBucket([]byte("MyBucket"))
	if err3 != nil {
		fmt.Println("Erreur création bucket")
	}

	err4 := b.Put([]byte("testclé"), []byte("testvaleur"))
	if err4 != nil {
		fmt.Println("Erreur put")
	}

	br := tx.Bucket([]byte("MyBucket"))
	v := br.Get([]byte("testclé"))
	fmt.Printf("Le résultat serait : %s\n", v)

}
