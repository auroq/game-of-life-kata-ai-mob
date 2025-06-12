package main

type Cell struct {
	alive bool
}

type Grid struct {
	cells  [][]Cell
	width  int
	height int
}

func (c Cell) nextState(neighbors int) bool {
	if c.alive {
		// Live cell rules
		if neighbors < 2 {
			return false // underpopulation
		}
		if neighbors == 2 || neighbors == 3 {
			return true // survival
		}
		return false // overpopulation
	} else {
		// Dead cell rules
		if neighbors == 3 {
			return true // reproduction
		}
		return false // stays dead
	}
}

func NewGrid(width, height int) *Grid {
	cells := make([][]Cell, height)
	for i := range cells {
		cells[i] = make([]Cell, width)
	}
	return &Grid{
		cells:  cells,
		width:  width,
		height: height,
	}
}

func (g *Grid) SetCell(x, y int, cell Cell) {
	g.cells[y][x] = cell
}

func (g *Grid) CountNeighbors(x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue // skip the center cell itself
			}
			
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < g.width && ny >= 0 && ny < g.height {
				if g.cells[ny][nx].alive {
					count++
				}
			}
		}
	}
	return count
}

func (g *Grid) GetCell(x, y int) Cell {
	return g.cells[y][x]
}

func (g *Grid) NextGeneration() *Grid {
	nextGrid := NewGrid(g.width, g.height)
	
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			currentCell := g.GetCell(x, y)
			neighbors := g.CountNeighbors(x, y)
			newAlive := currentCell.nextState(neighbors)
			nextGrid.SetCell(x, y, Cell{alive: newAlive})
		}
	}
	
	return nextGrid
}