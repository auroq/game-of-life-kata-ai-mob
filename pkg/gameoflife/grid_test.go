package gameoflife

import "testing"

func TestGrid(t *testing.T) {
	t.Run("when counting neighbors", func(t *testing.T) {
		t.Run("and center cell has no live neighbors", func(t *testing.T) {
			t.Run("it should return 0", func(t *testing.T) {
				grid := NewGrid(3, 3)
				neighbors := grid.CountNeighbors(1, 1)
				if neighbors != 0 {
					t.Errorf("Expected 0 neighbors, got %d", neighbors)
				}
			})
		})
		
		t.Run("and center cell has one live neighbor", func(t *testing.T) {
			t.Run("it should return 1", func(t *testing.T) {
				grid := NewGrid(3, 3)
				grid.SetCell(0, 0, true) // top-left neighbor
				neighbors := grid.CountNeighbors(1, 1)
				if neighbors != 1 {
					t.Errorf("Expected 1 neighbor, got %d", neighbors)
				}
			})
		})
		
		t.Run("and center cell has three live neighbors", func(t *testing.T) {
			t.Run("it should return 3", func(t *testing.T) {
				grid := NewGrid(3, 3)
				grid.SetCell(0, 0, true) // top-left
				grid.SetCell(1, 0, true) // top-center
				grid.SetCell(2, 0, true) // top-right
				neighbors := grid.CountNeighbors(1, 1)
				if neighbors != 3 {
					t.Errorf("Expected 3 neighbors, got %d", neighbors)
				}
			})
		})
	})
	
	t.Run("when evolving to next generation", func(t *testing.T) {
		t.Run("and dead cell has exactly 3 neighbors", func(t *testing.T) {
			t.Run("it should become alive", func(t *testing.T) {
				grid := NewGrid(3, 3)
				// Set up 3 neighbors around center (1,1)
				grid.SetCell(0, 0, true) // top-left
				grid.SetCell(1, 0, true) // top-center  
				grid.SetCell(2, 0, true) // top-right
				// Center cell (1,1) is dead by default
				
				nextGrid := grid.NextGeneration()
				
				if !nextGrid.GetCell(1, 1) {
					t.Error("Expected dead cell with 3 neighbors to become alive")
				}
			})
		})
		
		t.Run("and live cell has 2 neighbors", func(t *testing.T) {
			t.Run("it should survive", func(t *testing.T) {
				grid := NewGrid(3, 3)
				grid.SetCell(1, 1, true) // center cell alive
				grid.SetCell(0, 0, true) // top-left neighbor
				grid.SetCell(2, 2, true) // bottom-right neighbor
				
				nextGrid := grid.NextGeneration()
				
				if !nextGrid.GetCell(1, 1) {
					t.Error("Expected live cell with 2 neighbors to survive")
				}
			})
		})
		
		t.Run("and live cell has 4 neighbors", func(t *testing.T) {
			t.Run("it should die from overpopulation", func(t *testing.T) {
				grid := NewGrid(3, 3)
				grid.SetCell(1, 1, true) // center cell alive
				grid.SetCell(0, 0, true) // top-left
				grid.SetCell(1, 0, true) // top-center
				grid.SetCell(2, 0, true) // top-right
				grid.SetCell(0, 1, true) // left
				
				nextGrid := grid.NextGeneration()
				
				if nextGrid.GetCell(1, 1) {
					t.Error("Expected live cell with 4 neighbors to die from overpopulation")
				}
			})
		})
		
		t.Run("and live cell has 1 neighbor", func(t *testing.T) {
			t.Run("it should die from underpopulation", func(t *testing.T) {
				grid := NewGrid(3, 3)
				grid.SetCell(1, 1, true) // center cell alive
				grid.SetCell(0, 0, true) // only one neighbor
				
				nextGrid := grid.NextGeneration()
				
				if nextGrid.GetCell(1, 1) {
					t.Error("Expected live cell with 1 neighbor to die from underpopulation")
				}
			})
		})
	})
}