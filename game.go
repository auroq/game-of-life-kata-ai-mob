package main

type Grid struct {
	cells  [][]bool
	width  int
	height int
}

func NewGrid(width, height int) *Grid {
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}
	return &Grid{
		cells:  cells,
		width:  width,
		height: height,
	}
}

func (g *Grid) SetCell(x, y int, state bool) {
	g.cells[y][x] = state
}

func (g *Grid) GetCell(x, y int) bool {
	return g.cells[y][x]
}

func (g *Grid) CountNeighbors(x, y int) int {
	count := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue // Skip the cell itself
			}
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < g.width && ny >= 0 && ny < g.height {
				if g.cells[ny][nx] {
					count++
				}
			}
		}
	}
	return count
}

func (g *Grid) NextGeneration() *Grid {
	nextGrid := NewGrid(g.width, g.height)
	
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			currentState := g.GetCell(x, y)
			neighborCount := g.CountNeighbors(x, y)
			nextState := CellNextState(currentState, neighborCount)
			nextGrid.SetCell(x, y, nextState)
		}
	}
	
	return nextGrid
}

func CellNextState(currentState bool, neighborCount int) bool {
	if currentState {
		// Live cell
		if neighborCount == 2 || neighborCount == 3 {
			return true  // Survives
		}
		return false  // Dies (underpopulation or overpopulation)
	} else {
		// Dead cell
		if neighborCount == 3 {
			return true  // Reproduction
		}
		return false  // Stays dead
	}
}