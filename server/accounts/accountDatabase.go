package accounts

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
)

var db *bolt.DB
var err error

// Permet l'ouverture de la base de données
func OpenDB() {
	db, err = bolt.Open("bourbaki.db", 0600, nil)
	if err != nil {
		fmt.Print("Erreur à l'ouverture de la base de données")
	}
}

// Ajoute a dans la db. Retourne vrai si ça a fonctionné, faux sinon
func addInDB(cle string, a Account) bool {
	astring, _ := json.Marshal(a)
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte("Accounts"))
		err = b.Put([]byte(cle), []byte(astring))
		return err
	})
	if err != nil {
		fmt.Println("Erreur update")
		return false
	}
	return true
}

// Récupère l'account ayant pour clé celle passée en paramètre
func getFromDB(cle string) (res Account) {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Accounts"))
		v := b.Get([]byte(cle))
		json.Unmarshal(v, &res)
		return nil
	})
	return
}

// Permet de faire un test complet de toutes les fonctions, ajout et suppression
func Testsql() {
	OpenDB()
	a := Account{"yeti", "desneiges", 12}
	fonctionne := addInDB(a.Name, a)
	fmt.Println(fonctionne)

	var res Account
	res = getFromDB("cheval")
	fmt.Println(res.Name)
}
