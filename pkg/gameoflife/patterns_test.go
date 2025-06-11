package gameoflife

import "testing"

func TestClassicPatterns(t *testing.T) {
	t.Run("when testing blinker pattern", func(t *testing.T) {
		t.Run("and horizontal blinker evolves", func(t *testing.T) {
			t.Run("it should become vertical", func(t *testing.T) {
				// Create 5x5 grid for blinker pattern
				grid := NewGrid(5, 5)
				
				// Set up horizontal blinker: XXX
				grid.SetCell(1, 2, true) // left
				grid.SetCell(2, 2, true) // center  
				grid.SetCell(3, 2, true) // right
				
				nextGrid := grid.NextGeneration()
				
				// Should become vertical blinker
				if !nextGrid.GetCell(2, 1) || !nextGrid.GetCell(2, 2) || !nextGrid.GetCell(2, 3) {
					t.Error("Expected horizontal blinker to become vertical")
				}
				
				// Original horizontal positions should be dead
				if nextGrid.GetCell(1, 2) || nextGrid.GetCell(3, 2) {
					t.Error("Expected horizontal blinker positions to be dead after evolution")
				}
			})
		})
		
		t.Run("and vertical blinker evolves", func(t *testing.T) {
			t.Run("it should become horizontal", func(t *testing.T) {
				grid := NewGrid(5, 5)
				
				// Set up vertical blinker
				grid.SetCell(2, 1, true) // top
				grid.SetCell(2, 2, true) // center
				grid.SetCell(2, 3, true) // bottom
				
				nextGrid := grid.NextGeneration()
				
				// Should become horizontal blinker
				if !nextGrid.GetCell(1, 2) || !nextGrid.GetCell(2, 2) || !nextGrid.GetCell(3, 2) {
					t.Error("Expected vertical blinker to become horizontal")
				}
				
				// Original vertical positions should be dead (except center)
				if nextGrid.GetCell(2, 1) || nextGrid.GetCell(2, 3) {
					t.Error("Expected vertical blinker end positions to be dead after evolution")
				}
			})
		})
	})
	
	t.Run("when testing block pattern", func(t *testing.T) {
		t.Run("and block evolves", func(t *testing.T) {
			t.Run("it should remain stable", func(t *testing.T) {
				grid := NewGrid(4, 4)
				
				// Set up 2x2 block pattern
				grid.SetCell(1, 1, true) // top-left
				grid.SetCell(2, 1, true) // top-right
				grid.SetCell(1, 2, true) // bottom-left
				grid.SetCell(2, 2, true) // bottom-right
				
				nextGrid := grid.NextGeneration()
				
				// Block should remain stable
				if !nextGrid.GetCell(1, 1) || !nextGrid.GetCell(2, 1) || 
				   !nextGrid.GetCell(1, 2) || !nextGrid.GetCell(2, 2) {
					t.Error("Expected block pattern to remain stable")
				}
				
				// Verify no other cells are alive
				for y := 0; y < 4; y++ {
					for x := 0; x < 4; x++ {
						if (x == 1 || x == 2) && (y == 1 || y == 2) {
							continue // Skip the block cells
						}
						if nextGrid.GetCell(x, y) {
							t.Errorf("Expected cell (%d,%d) to be dead in stable block", x, y)
						}
					}
				}
			})
		})
	})
}