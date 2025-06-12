package main

func evolveCell(isAlive bool, neighbors int) bool {
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
	
	for i := row - 1; i <= row + 1; i++ {
		for j := col - 1; j <= col + 1; j++ {
			if i == row && j == col {
				continue // skip the cell itself
			}
			if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) {
				if grid[i][j] {
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
	
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			neighbors := countNeighbors(grid, i, j)
			newGrid[i][j] = evolveCell(grid[i][j], neighbors)
		}
	}
	
	return newGrid
}

func gridsEqual(grid1, grid2 [][]bool) bool {
	if len(grid1) != len(grid2) {
		return false
	}
	
	for i := range grid1 {
		if len(grid1[i]) != len(grid2[i]) {
			return false
		}
		for j := range grid1[i] {
			if grid1[i][j] != grid2[i][j] {
				return false
			}
		}
	}
	
	return true
}