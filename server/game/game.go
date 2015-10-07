package game

import(
  "go-bourbaki/server/globals"
)
// Game structure définissant une partie
type Game struct {
  lines [globals.GRIDSIZE][globals.GRIDSIZE][2]int
  squares [globals.GRIDSIZE][globals.GRIDSIZE]int

}
