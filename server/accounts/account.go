package accounts

import (
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

// Login Fonction qui connecte l'utilisateur ou qui lui crée un compte s'il n'en a pas
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

// IsValidUsername Retourne true si le pseudo choisi est correct (longueur, caractères spéciaux...)
func IsValidUsername(name string) bool {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9-_]{3,15}$", name)
	if err != nil {
		globals.ErrLogger.Println("Erreur sur la regex sur " + name)
		return false
	}
	return matched
}

// GetGeneralLadder ..
func GetGeneralLadder() globals.Classement{
	var classementtb globals.Classement = GetAllAccounts()
	// 1e étape: récupération du classement
	for _, player := range classementtb {
		p := globals.PlayerClassement{0, player.NumPlayer, player.Name, player.Score, player.NbrGames, player.NbrWins}
		classementtb = append(classementtb, p)
	}

	// 2e étape: tri par Score
	sort.Sort(globals.ByScore{classementtb})

	// 3e étape: ajout de l'attribut Classement
	for i := 1; i <= len(classementtb); i++ {
		classementtb[i-1].Classement = i
	}
	accountLogger.Println("Classement: ", classementtb)

	return classementtb
}
