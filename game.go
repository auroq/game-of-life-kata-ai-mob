package main

type Cell struct {
	Alive bool
}

type Grid struct {
	cells [][]Cell
}

func (c Cell) NextState(neighbors int) Cell {
	if !c.Alive && neighbors == 3 {
		return Cell{Alive: true}
	}
	if c.Alive && (neighbors == 2 || neighbors == 3) {
		return Cell{Alive: true}
	}
	return Cell{Alive: false}
}

func (g Grid) CountNeighbors(row, col int) int {
	count := 0
	for r := row - 1; r <= row + 1; r++ {
		for c := col - 1; c <= col + 1; c++ {
			if r == row && c == col {
				continue
			}
			if r >= 0 && r < len(g.cells) && c >= 0 && c < len(g.cells[0]) {
				if g.cells[r][c].Alive {
					count++
				}
			}
		}
	}
	return count
}

func (g Grid) NextGeneration() Grid {
	nextCells := make([][]Cell, len(g.cells))
	for r := range g.cells {
		nextCells[r] = make([]Cell, len(g.cells[r]))
		for c := range g.cells[r] {
			neighbors := g.CountNeighbors(r, c)
			nextCells[r][c] = g.cells[r][c].NextState(neighbors)
		}
	}
	return Grid{cells: nextCells}
}