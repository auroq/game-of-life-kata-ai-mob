package main

func nextState(isAlive bool, neighbors int) bool {
	if !isAlive && neighbors == 3 {
		return true
	}
	if isAlive && (neighbors == 2 || neighbors == 3) {
		return true
	}
	return false
}

func countNeighbors(grid [][]bool, row, col int) int {
	count := 0
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			r, c := row+dr, col+dc
			if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) {
				if grid[r][c] {
					count++
				}
			}
		}
	}
	return count
}

func evolveGrid(grid [][]bool) [][]bool {
	rows := len(grid)
	cols := len(grid[0])
	newGrid := make([][]bool, rows)
	for i := range newGrid {
		newGrid[i] = make([]bool, cols)
	}
	
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			neighbors := countNeighbors(grid, r, c)
			newGrid[r][c] = nextState(grid[r][c], neighbors)
		}
	}
	
	return newGrid
}