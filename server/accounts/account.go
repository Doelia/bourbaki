package accounts

type Account struct {
	Name, Pass string // TODO penser à md5 le mdp
	Points     int
}
