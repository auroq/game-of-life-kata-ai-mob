package main

import (
	"testing"
)

func TestCellNextState(t *testing.T) {
	t.Run("when cell is alive", func(t *testing.T) {
		t.Run("and has zero neighbors", func(t *testing.T) {
			t.Run("it should die", func(t *testing.T) {
				currentState := true  // alive
				neighborCount := 0
				
				nextState := CellNextState(currentState, neighborCount)
				
				if nextState != false {
					t.Errorf("Expected dead cell (false), got %v", nextState)
				}
			})
		})
		
		t.Run("and has two neighbors", func(t *testing.T) {
			t.Run("it should survive", func(t *testing.T) {
				currentState := true  // alive
				neighborCount := 2
				
				nextState := CellNextState(currentState, neighborCount)
				
				if nextState != true {
					t.Errorf("Expected live cell (true), got %v", nextState)
				}
			})
		})
		
		t.Run("and has three neighbors", func(t *testing.T) {
			t.Run("it should survive", func(t *testing.T) {
				currentState := true  // alive
				neighborCount := 3
				
				nextState := CellNextState(currentState, neighborCount)
				
				if nextState != true {
					t.Errorf("Expected live cell (true), got %v", nextState)
				}
			})
		})
		
		t.Run("and has four neighbors", func(t *testing.T) {
			t.Run("it should die", func(t *testing.T) {
				currentState := true  // alive
				neighborCount := 4
				
				nextState := CellNextState(currentState, neighborCount)
				
				if nextState != false {
					t.Errorf("Expected dead cell (false), got %v", nextState)
				}
			})
		})
	})
	
	t.Run("when cell is dead", func(t *testing.T) {
		t.Run("and has three neighbors", func(t *testing.T) {
			t.Run("it should become alive", func(t *testing.T) {
				currentState := false  // dead
				neighborCount := 3
				
				nextState := CellNextState(currentState, neighborCount)
				
				if nextState != true {
					t.Errorf("Expected live cell (true), got %v", nextState)
				}
			})
		})
		
		t.Run("and has two neighbors", func(t *testing.T) {
			t.Run("it should stay dead", func(t *testing.T) {
				currentState := false  // dead
				neighborCount := 2
				
				nextState := CellNextState(currentState, neighborCount)
				
				if nextState != false {
					t.Errorf("Expected dead cell (false), got %v", nextState)
				}
			})
		})
	})
}

func TestGrid(t *testing.T) {
	t.Run("when creating a new grid", func(t *testing.T) {
		t.Run("and setting a cell to alive", func(t *testing.T) {
			t.Run("it should return alive when getting that cell", func(t *testing.T) {
				grid := NewGrid(3, 3)
				
				grid.SetCell(1, 1, true)
				cellState := grid.GetCell(1, 1)
				
				if cellState != true {
					t.Errorf("Expected cell to be alive (true), got %v", cellState)
				}
			})
		})
		
		t.Run("and accessing an unset cell", func(t *testing.T) {
			t.Run("it should be dead by default", func(t *testing.T) {
				grid := NewGrid(3, 3)
				
				cellState := grid.GetCell(0, 0)
				
				if cellState != false {
					t.Errorf("Expected cell to be dead (false), got %v", cellState)
				}
			})
		})
	})
}

func TestNeighborCounting(t *testing.T) {
	t.Run("when cell has one live neighbor", func(t *testing.T) {
		t.Run("it should count exactly one neighbor", func(t *testing.T) {
			grid := NewGrid(3, 3)
			
			grid.SetCell(1, 1, true)  // Set center cell alive
			neighborCount := grid.CountNeighbors(0, 0)  // Count neighbors of top-left cell
			
			if neighborCount != 1 {
				t.Errorf("Expected 1 neighbor, got %d", neighborCount)
			}
		})
	})
	
	t.Run("when cell has three live neighbors", func(t *testing.T) {
		t.Run("it should count exactly three neighbors", func(t *testing.T) {
			grid := NewGrid(3, 3)
			
			// Set up pattern with center cell having 3 neighbors
			grid.SetCell(0, 1, true)  // Left neighbor
			grid.SetCell(1, 0, true)  // Top neighbor  
			grid.SetCell(2, 1, true)  // Right neighbor
			
			neighborCount := grid.CountNeighbors(1, 1)  // Count neighbors of center cell
			
			if neighborCount != 3 {
				t.Errorf("Expected 3 neighbors, got %d", neighborCount)
			}
		})
	})
}

func TestGameIteration(t *testing.T) {
	t.Run("when blinker pattern", func(t *testing.T) {
		t.Run("and iterating once", func(t *testing.T) {
			t.Run("it should rotate 90 degrees", func(t *testing.T) {
				grid := NewGrid(5, 5)
				
				// Set up vertical blinker
				grid.SetCell(2, 1, true)
				grid.SetCell(2, 2, true)
				grid.SetCell(2, 3, true)
				
				nextGrid := grid.NextGeneration()
				
				// Should become horizontal blinker
				if !nextGrid.GetCell(1, 2) || !nextGrid.GetCell(2, 2) || !nextGrid.GetCell(3, 2) {
					t.Error("Expected horizontal blinker pattern")
				}
				
				// Original vertical cells should be dead (except center)
				if nextGrid.GetCell(2, 1) || nextGrid.GetCell(2, 3) {
					t.Error("Expected vertical cells to die")
				}
			})
		})
	})
}