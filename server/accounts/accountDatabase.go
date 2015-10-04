package accounts

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"crypto/md5"
)

var db *bolt.DB
var err error

// OpenDB Permet l'ouverture de la base de données
func OpenDB() {
	db, err = bolt.Open("bourbaki.db", 0600, nil)
	if err != nil {
		fmt.Print("Erreur à l'ouverture de la base de données")
	}
}

// Permet l'ajout de a dans la bd
//@param cle string: Name de l'account à ajouter
//@param a Account : La structure account à ajouter
//@return bool : Vrai si l'ajout a bien été fait, faux sinon
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
//@param cle: Name de l'account à récupérer
//@return Account: Le compte correspondant
func getFromDB(cle string) (res Account) {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Accounts"))
		v := b.Get([]byte(cle))
		json.Unmarshal(v, &res)
		return nil
	})
	return
}

// Permet de savoir si un compte existe pour la clé ou pas
//@param cle: Name de l'account recherché
//@return Account: Le compte s'il existe
//@return bool: Vrai si le compte existe, faux sinon
func exists(cle string) (Account, bool) {
	var res Account
	db.View(func(tx *bolt.Tx) error {
		var v []byte
		b := tx.Bucket([]byte("Accounts"))
		v = b.Get([]byte(cle))
		json.Unmarshal(v, &res)
		return nil
	})
	return res, res.Name != ""
}

// CreateAccount Permet de créer une structure Account
//@param n: name de l'account
//@param p: mot de passe de l'account
func CreateAccount(n string, p string) Account{
	mdph := md5.Sum( []byte(p))
	ac := Account{n, mdph, 0}
	return ac
}

// Testsql Permet de faire un test complet de toutes les fonctions, ajout et suppression
func Testsql() {
	OpenDB()
	a := CreateAccount("yeti", "deseaux")
	fonctionne := addInDB(a.Name, a)
	fmt.Println(fonctionne)

	var res Account
	res = getFromDB("yeti")
	fmt.Println(res)
/*	res, b := exists("yeti")
	fmt.Println(res)
	fmt.Println(b)*/
	//il faut tester exists et getFromDB
}
