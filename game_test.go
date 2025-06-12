package main

import (
	"testing"
)

func TestGameOfLife(t *testing.T) {
	t.Run("when a live cell has fewer than 2 neighbors", func(t *testing.T) {
		t.Run("it should die", func(t *testing.T) {
			cell := Cell{alive: true}
			neighbors := 0
			
			result := cell.nextState(neighbors)
			
			if result != false {
				t.Errorf("Expected cell to die, but it survived")
			}
		})
	})
	
	t.Run("when a live cell has 2 neighbors", func(t *testing.T) {
		t.Run("it should survive", func(t *testing.T) {
			cell := Cell{alive: true}
			neighbors := 2
			
			result := cell.nextState(neighbors)
			
			if result != true {
				t.Errorf("Expected cell to survive, but it died")
			}
		})
	})
	
	t.Run("when a live cell has 3 neighbors", func(t *testing.T) {
		t.Run("it should survive", func(t *testing.T) {
			cell := Cell{alive: true}
			neighbors := 3
			
			result := cell.nextState(neighbors)
			
			if result != true {
				t.Errorf("Expected cell to survive, but it died")
			}
		})
	})
	
	t.Run("when a live cell has more than 3 neighbors", func(t *testing.T) {
		t.Run("it should die from overpopulation", func(t *testing.T) {
			cell := Cell{alive: true}
			neighbors := 4
			
			result := cell.nextState(neighbors)
			
			if result != false {
				t.Errorf("Expected cell to die from overpopulation, but it survived")
			}
		})
	})
	
	t.Run("when a dead cell has exactly 3 neighbors", func(t *testing.T) {
		t.Run("it should become alive by reproduction", func(t *testing.T) {
			cell := Cell{alive: false}
			neighbors := 3
			
			result := cell.nextState(neighbors)
			
			if result != true {
				t.Errorf("Expected dead cell to become alive, but it stayed dead")
			}
		})
	})
	
	t.Run("when a dead cell has 2 neighbors", func(t *testing.T) {
		t.Run("it should stay dead", func(t *testing.T) {
			cell := Cell{alive: false}
			neighbors := 2
			
			result := cell.nextState(neighbors)
			
			if result != false {
				t.Errorf("Expected dead cell to stay dead, but it became alive")
			}
		})
	})
}

func TestGrid(t *testing.T) {
	t.Run("when counting neighbors in a 3x3 grid", func(t *testing.T) {
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
				grid.SetCell(0, 0, Cell{alive: true})
				
				neighbors := grid.CountNeighbors(1, 1)
				
				if neighbors != 1 {
					t.Errorf("Expected 1 neighbor, got %d", neighbors)
				}
			})
		})
	})
}

func TestGridEvolution(t *testing.T) {
	t.Run("when evolving to next generation", func(t *testing.T) {
		t.Run("and a single live cell has no neighbors", func(t *testing.T) {
			t.Run("it should die", func(t *testing.T) {
				grid := NewGrid(3, 3)
				grid.SetCell(1, 1, Cell{alive: true})
				
				nextGrid := grid.NextGeneration()
				
				cell := nextGrid.GetCell(1, 1)
				if cell.alive {
					t.Errorf("Expected cell to die, but it was alive")
				}
			})
		})
		
		t.Run("and a stable block pattern is present", func(t *testing.T) {
			t.Run("it should remain unchanged", func(t *testing.T) {
				grid := NewGrid(4, 4)
				// Create 2x2 block pattern
				grid.SetCell(1, 1, Cell{alive: true})
				grid.SetCell(2, 1, Cell{alive: true})
				grid.SetCell(1, 2, Cell{alive: true})
				grid.SetCell(2, 2, Cell{alive: true})
				
				nextGrid := grid.NextGeneration()
				
				// All block cells should still be alive
				if !nextGrid.GetCell(1, 1).alive {
					t.Errorf("Expected block cell (1,1) to stay alive")
				}
				if !nextGrid.GetCell(2, 1).alive {
					t.Errorf("Expected block cell (2,1) to stay alive")
				}
				if !nextGrid.GetCell(1, 2).alive {
					t.Errorf("Expected block cell (1,2) to stay alive")
				}
				if !nextGrid.GetCell(2, 2).alive {
					t.Errorf("Expected block cell (2,2) to stay alive")
				}
			})
		})
	})
}