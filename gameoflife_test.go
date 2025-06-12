package main

import "testing"

func TestCellEvolution(t *testing.T) {
	t.Run("when a dead cell has no neighbors", func(t *testing.T) {
		t.Run("it should remain dead", func(t *testing.T) {
			cell := false // dead cell
			neighbors := 0
			
			result := evolveCell(cell, neighbors)
			
			if result != false {
				t.Errorf("Expected dead cell to remain dead, but got %v", result)
			}
		})
	})
	
	t.Run("when a dead cell has exactly 3 neighbors", func(t *testing.T) {
		t.Run("it should become alive", func(t *testing.T) {
			cell := false // dead cell
			neighbors := 3
			
			result := evolveCell(cell, neighbors)
			
			if result != true {
				t.Errorf("Expected dead cell with 3 neighbors to become alive, but got %v", result)
			}
		})
	})
	
	t.Run("when a live cell has 2 neighbors", func(t *testing.T) {
		t.Run("it should survive", func(t *testing.T) {
			cell := true // live cell
			neighbors := 2
			
			result := evolveCell(cell, neighbors)
			
			if result != true {
				t.Errorf("Expected live cell with 2 neighbors to survive, but got %v", result)
			}
		})
	})
	
	t.Run("when a live cell has 3 neighbors", func(t *testing.T) {
		t.Run("it should survive", func(t *testing.T) {
			cell := true // live cell
			neighbors := 3
			
			result := evolveCell(cell, neighbors)
			
			if result != true {
				t.Errorf("Expected live cell with 3 neighbors to survive, but got %v", result)
			}
		})
	})
	
	t.Run("when a live cell has 1 neighbor", func(t *testing.T) {
		t.Run("it should die from underpopulation", func(t *testing.T) {
			cell := true // live cell
			neighbors := 1
			
			result := evolveCell(cell, neighbors)
			
			if result != false {
				t.Errorf("Expected live cell with 1 neighbor to die, but got %v", result)
			}
		})
	})
	
	t.Run("when a live cell has 4 neighbors", func(t *testing.T) {
		t.Run("it should die from overpopulation", func(t *testing.T) {
			cell := true // live cell
			neighbors := 4
			
			result := evolveCell(cell, neighbors)
			
			if result != false {
				t.Errorf("Expected live cell with 4 neighbors to die, but got %v", result)
			}
		})
	})
}

func TestNeighborCounting(t *testing.T) {
	t.Run("when counting neighbors in a 3x3 grid", func(t *testing.T) {
		t.Run("and center cell has no live neighbors", func(t *testing.T) {
			t.Run("it should return 0", func(t *testing.T) {
				grid := [][]bool{
					{false, false, false},
					{false, false, false},
					{false, false, false},
				}
				
				count := countNeighbors(grid, 1, 1)
				
				if count != 0 {
					t.Errorf("Expected 0 neighbors, but got %d", count)
				}
			})
		})
		
		t.Run("and center cell has 1 live neighbor", func(t *testing.T) {
			t.Run("it should return 1", func(t *testing.T) {
				grid := [][]bool{
					{true, false, false},
					{false, false, false},
					{false, false, false},
				}
				
				count := countNeighbors(grid, 1, 1)
				
				if count != 1 {
					t.Errorf("Expected 1 neighbor, but got %d", count)
				}
			})
		})
		
		t.Run("and center cell has 3 live neighbors", func(t *testing.T) {
			t.Run("it should return 3", func(t *testing.T) {
				grid := [][]bool{
					{true, true, false},
					{true, false, false},
					{false, false, false},
				}
				
				count := countNeighbors(grid, 1, 1)
				
				if count != 3 {
					t.Errorf("Expected 3 neighbors, but got %d", count)
				}
			})
		})
	})
}

func TestGridEvolution(t *testing.T) {
	t.Run("when evolving a grid for one generation", func(t *testing.T) {
		t.Run("and a dead center cell has exactly 3 neighbors", func(t *testing.T) {
			t.Run("it should become alive", func(t *testing.T) {
				initialGrid := [][]bool{
					{true, true, false},
					{true, false, false},
					{false, false, false},
				}
				
				expectedGrid := [][]bool{
					{true, true, false},
					{true, true, false},
					{false, false, false},
				}
				
				result := evolveGrid(initialGrid)
				
				if !gridsEqual(result, expectedGrid) {
					t.Errorf("Grid evolution failed. Expected center cell to become alive")
				}
			})
		})
		
		t.Run("and testing the blinker oscillator pattern", func(t *testing.T) {
			t.Run("it should oscillate between horizontal and vertical", func(t *testing.T) {
				horizontalBlinker := [][]bool{
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, true, true, true, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
				}
				
				verticalBlinker := [][]bool{
					{false, false, false, false, false},
					{false, false, true, false, false},
					{false, false, true, false, false},
					{false, false, true, false, false},
					{false, false, false, false, false},
				}
				
				generation1 := evolveGrid(horizontalBlinker)
				generation2 := evolveGrid(generation1)
				
				if !gridsEqual(generation1, verticalBlinker) {
					t.Errorf("First generation should be vertical blinker")
				}
				
				if !gridsEqual(generation2, horizontalBlinker) {
					t.Errorf("Second generation should return to horizontal blinker")
				}
			})
		})
	})
}