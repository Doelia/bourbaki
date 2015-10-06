package accounts

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"go-bourbaki/server/globals"
	"github.com/boltdb/bolt"
)

var db *bolt.DB
var err error

// OpenDB Permet l'ouverture de la base de données
func OpenDB() {
	db, err = bolt.Open("bourbaki.db", 0600, nil)
	if err != nil {
		globals.ErrLogger.Println("Erreur à l'ouverture de la base de données")
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Accounts"))
		return err
	})
}

// Permet l'ajout de a dans la bd
//@param cle string: Name de l'account à ajouter
//@param a Account : La structure account à ajouter
//@return bool : Vrai si l'ajout a bien été fait, faux sinon
func addInDB(cle string, account Account) bool {
	jsonaccount, _ := json.Marshal(account)
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Accounts"))
		err = bucket.Put([]byte(cle), []byte(jsonaccount))
		return err
	})
	if err != nil {
		globals.ErrLogger.Println("Erreur update")
		return false
	}
	return true
}

// Récupère l'account ayant pour clé celle passée en paramètre
//@param cle: Name de l'account à récupérer
//@return Account: Le compte correspondant
func getFromDB(cle string) (account Account) {
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Accounts"))
		v := bucket.Get([]byte(cle))
		json.Unmarshal(v, &account)
		return nil
	})
	return
}

// Exists Permet de savoir si un compte existe pour la clé ou pas
//@param cle: Name de l'account recherché
//@return Account: Le compte s'il existe
//@return bool: Vrai si le compte existe, faux sinon
func Exists(cle string) (Account, bool) {
	var account Account
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Accounts"))
		v := bucket.Get([]byte(cle))
		json.Unmarshal(v, &account)
		return nil
	})
	return account, account.Name != ""
}

// CreateAccount Permet de créer une structure Account
//@param n: name de l'account
//@param p: mot de passe de l'account
func CreateAccount(name string, pass string) Account {
	motdepassemd5 := md5.Sum([]byte(pass))
	account := Account{name, motdepassemd5, 0}
	return account
}

// Testsql Permet de faire un test complet de toutes les fonctions, ajout et suppression
func Testsql() {
	testaccount1 := CreateAccount("anne", "motdepasseanne")
	fmt.Println(addInDB(testaccount1.Name, testaccount1))

	testaccount2 := CreateAccount("henri", "motdepassehenri")
	fmt.Println(addInDB(testaccount2.Name, testaccount2))

	// Test de getFromDB
	var testaccountget Account
	testaccountget = getFromDB("yeti")
	fmt.Println(testaccountget)

	// Test de exists
	testaccountexist, resultatboolexist := Exists("yeti")
	fmt.Println(testaccountexist)
	fmt.Println(resultatboolexist)

	_, r := Login("anne", "motdepasseanne")
	fmt.Println(r) // doit retourner 1
	_, r = Login("anne", "motdepassehenri")
	fmt.Println(r) // doit retourner 0
	_, r = Login("caly", "motdepassecaly")
	fmt.Println(r) // doit retourner 2
}
