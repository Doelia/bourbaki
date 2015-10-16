package game

import "go-bourbaki/server/globals"

// isActive Retourne vrai si la ligne est active dans le game, faux sinon
func (g *Game) isActive(x int, y int, o int) bool {
	if x < 0 || x >= globals.GRIDSIZE {
		return false
	}
	if y < 0 || y >= globals.GRIDSIZE {
		return false
	}
	return g.lines[x][y][o] > 0
}

// IsPlayable Retourne vrai si la ligne est jouable, faux sinon (élimination cas de triche)
func (g *Game) IsPlayable(x int, y int, o int) bool {
	if x < 0 || x >= globals.GRIDSIZE {
		return false
	}
	if y < 0 || y >= globals.GRIDSIZE {
		return false
	}
	return g.lines[x][y][o] == 0
}

// AddLine Active la ligne dans le game
func (g *Game) AddLine(line globals.Line) {
	if g.lines[line.X][line.Y][line.O] == 0 {
		g.lines[line.X][line.Y][line.O] = line.N
	}
}

// AddSquare Active le carré dans le game
func (g *Game) AddSquare(square globals.Square) {
	g.squares[square.X][square.Y] = square.N
}

// TestSquare permet de savoir si la ligne qui vient d'être jouée forme un carré
// @param lastLine: dernière ligne ayant été jouée
// @return bool: vrai si le joueur gagne un carré, faux sinon
// @return list: liste de tous les carrés formés
func (g *Game) TestSquare(lastLine globals.Line) (isSquare bool, list []globals.Square) {
	x := lastLine.X
	y := lastLine.Y
	isSquare = false
	if lastLine.O == globals.HORIZONTAL {
		if g.isActive(x, y-1, globals.HORIZONTAL) && g.isActive(x+1, y-1, globals.VERTICAL) && g.isActive(x, y-1, globals.VERTICAL) {
			isSquare = true
			list = append(list, globals.Square{x, y - 1, lastLine.N})
		}
		if g.isActive(x, y+1, globals.HORIZONTAL) && g.isActive(x, y, globals.VERTICAL) && g.isActive(x+1, y, globals.VERTICAL) {
			isSquare = true
			list = append(list, globals.Square{x, y, lastLine.N})
		}
	} else {
		if lastLine.O == globals.VERTICAL {
			if g.isActive(x, y, globals.HORIZONTAL) && g.isActive(x+1, y, globals.VERTICAL) && g.isActive(x, y+1, globals.HORIZONTAL) {
				isSquare = true
				list = append(list, globals.Square{x, y, lastLine.N})
			}
			if g.isActive(x-1, y, 0) && g.isActive(x-1, y, 1) && g.isActive(x-1, y+1, 0) {
				isSquare = true
				list = append(list, globals.Square{x - 1, y, lastLine.N})
			}
		}
	}
	return
}

// GetActivesLinesList Retourne une liste de toutes les lignes actives (déjà posées)
func (g *Game) GetActivesLinesList() (list []globals.Line) {
	for i := 0; i < globals.GRIDSIZE; i++ {
		for j := 0; j < globals.GRIDSIZE; j++ {
			for o := 0; o <= 1; o++ {
				if g.lines[i][j][o] > 0 {
					list = append(list, globals.Line{i, j, o, g.lines[i][j][o]})
				}
			}
		}
	}
	return
}

// GetActivesSquaresList Retourne une liste de tous les carrés remplis
func (g *Game) GetActivesSquaresList() (list []globals.Square) {
	for i := 0; i < globals.GRIDSIZE-1; i++ {
		for j := 0; j < globals.GRIDSIZE-1; j++ {
			if g.squares[i][j] > 0 {
				list = append(list, globals.Square{i, j, g.squares[i][j]})
			}
		}
	}
	return
}
