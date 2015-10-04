package accounts

// Account structure
type Account struct {
	Name   string
	Pass   [16]byte // TODO penser à md5 le mdp
	Points int
}

//Fonction qui connecte l'utilisateur ou qui lui crée un compte s'il n'en a pas
func Login(name string, pass string) (Account, int) {
	//cas 1: bon name et pass
	OpenDB()
	a := CreateAccount(name, pass)
	res, b := Exists(name)
	if b == true {
		if res.Pass == a.Pass {
			// le compte existe déjà, connexion réussie
			return a, 1
		} else {
			// le compte existe déjà mais le mot de passe est éronné
			return a, 0
		}
	} else {
		// ajout du nouveau compte dans la bd, connexion réussie
		fonctionne := addInDB(a.Name, a)
		if fonctionne == true {
			return a, 2
		} else {
			return a, 3 // erreur interne
		}
	}
}
