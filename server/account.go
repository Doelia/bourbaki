package main

//Penser à la visibilité
type Account struct {
	Name, Pass string //penser à md5 le mdp
	Points     int
}

//permet d'ajouter le compte dans la BD
func (a Account) AddAccountinDB() bool {
	//	result, err := db.Exec("INSERT INTO Accounts VALUES (?, ?, ?)", a.Name, a.Pass, a.Points)
	/*	if err != nil {
		log.Fatal(err)
		return false
	}*/
	return true
}
