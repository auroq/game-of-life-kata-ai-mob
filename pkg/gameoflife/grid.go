package gameoflife

type Grid struct {
	cells  [][]*Cell
	width  int
	height int
}

func NewGrid(width, height int) *Grid {
	cells := make([][]*Cell, height)
	for i := range cells {
		cells[i] = make([]*Cell, width)
		for j := range cells[i] {
			cells[i][j] = NewCell()
		}
	}
	return &Grid{
		cells:  cells,
		width:  width,
		height: height,
	}
}

func (g *Grid) SetCell(x, y int, alive bool) {
	g.cells[y][x].SetAlive(alive)
}

func (g *Grid) CountNeighbors(x, y int) int {
	count := 0
	
	// Check all 8 directions around the cell
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			// Skip the center cell itself
			if dx == 0 && dy == 0 {
				continue
			}
			
			nx, ny := x+dx, y+dy
			
			// Check bounds
			if nx >= 0 && nx < g.width && ny >= 0 && ny < g.height {
				if g.cells[ny][nx].IsAlive() {
					count++
				}
			}
		}
	}
	
	return count
}

func (g *Grid) GetCell(x, y int) bool {
	return g.cells[y][x].IsAlive()
}

func (g *Grid) NextGeneration() *Grid {
	nextGrid := NewGrid(g.width, g.height)
	
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			neighbors := g.CountNeighbors(x, y)
			currentlyAlive := g.GetCell(x, y)
			
			// Apply Game of Life rules
			if !currentlyAlive && neighbors == 3 {
				// Dead cell with exactly 3 neighbors becomes alive (reproduction)
				nextGrid.SetCell(x, y, true)
			} else if currentlyAlive && (neighbors == 2 || neighbors == 3) {
				// Live cell with 2 or 3 neighbors survives
				nextGrid.SetCell(x, y, true)
			}
			// All other cases: cell dies or stays dead
		}
	}
	
	return nextGrid
}

func (g *Grid) String() string {
	result := ""
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			if g.GetCell(x, y) {
				result += "█"
			} else {
				result += "."
			}
		}
		result += "\n"
	}
	return result
}