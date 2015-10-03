package accounts

// Account structure
type Account struct {
	Name, Pass string // TODO penser à md5 le mdp
	Points     int
}

//Fonction qui connecte l'utilisateur ou qui lui crée un compte s'il n'en a pas
//func login(string name, string  pass) {
//}
