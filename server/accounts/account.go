package accounts

import (
	"crypto/md5"
	"go-bourbaki/server/globals"
	"regexp"
	"sort"
)

// Account structure
type Account struct {
	Name     string   // Nom du joueur
	Pass     [16]byte // Mot de passe crypté
	Points   int      // Nombore de points accumulés
	NbrGames int      // Nombre de parties jouées
	NbrWins  int      // Nombre de parties gagnées
}

// CreateAccount Permet de créer une structure Account, avec le mot de passe crypté
// @param n: name de l'account
// @param p: mot de passe de l'account (non crypté)
func CreateAccount(name string, pass string) Account {
	motdepassemd5 := md5.Sum([]byte(pass))
	account := Account{name, motdepassemd5, 0, 0, 0}
	return account
}

// IsValidUsername Retourne true si le pseudo choisi est correct (longueur, caractères spéciaux...)
func IsValidUsername(name string) bool {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9-_]{3,15}$", name)
	if err != nil {
		globals.ErrLogger.Println("Erreur sur la regex sur " + name)
		return false
	}
	return matched
}

// Login Connecte l'utilisateur ou lui crée un compte s'il n'en a pas
// @param name: nom de l'account
// @param pass: mot de passe non crypté
// @return Account: le compte correspondant à l'utilisateur connecté
// @return int, 0 si mot de passe incorrect, 1 si connexion OK, 2 si connexion OK + compte créé
func Login(name string, pass string) (Account, int) {
	account := CreateAccount(name, pass)
	accountexistant, isExist := Exists(name)
	if isExist {
		if account.Pass == accountexistant.Pass {
			return account, 1
		}
		return account, 0
	}
	if addInDB(account.Name, account) {
		return account, 2
	}
	return account, 3 // erreur interne
}

// GetGeneralLadder Permet de récupérer le classement général
func GetGeneralLadder() globals.Classement {
	// 1e étape: récupération du classement dans la BD
	var classementaccount globals.Classement = GetAllAccounts()
	var classementplayer globals.Classement
	for _, player := range classementaccount {
		p := globals.PlayerClassement{0, player.NumPlayer, player.Name, player.Score, player.NbrGames, player.NbrWins}
		classementplayer = append(classementplayer, p)
	}

	// 2e étape: tri par Score
	sort.Sort(globals.ByScore{classementplayer})

	// 3e étape: ajout de l'attribut "Classement" correspondant à l'ordre des joueurs
	for i := 1; i <= len(classementplayer); i++ {
		classementplayer[i-1].Classement = i
	}
	accountLogger.Println("Classement: ", classementplayer)

	return classementplayer
}
