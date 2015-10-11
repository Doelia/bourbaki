package accounts

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
