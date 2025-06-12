package main

import "testing"

func TestCell(t *testing.T) {
	t.Run("when dead cell has no neighbors", func(t *testing.T) {
		t.Run("it should remain dead", func(t *testing.T) {
			cell := Cell{Alive: false}
			neighbors := 0
			
			result := cell.NextState(neighbors)
			
			if result.Alive != false {
				t.Errorf("Expected dead cell to remain dead, got alive")
			}
		})
	})
	
	t.Run("when dead cell has exactly 3 neighbors", func(t *testing.T) {
		t.Run("it should become alive", func(t *testing.T) {
			cell := Cell{Alive: false}
			neighbors := 3
			
			result := cell.NextState(neighbors)
			
			if result.Alive != true {
				t.Errorf("Expected dead cell with 3 neighbors to become alive, got dead")
			}
		})
	})
	
	t.Run("when live cell has 2 neighbors", func(t *testing.T) {
		t.Run("it should survive", func(t *testing.T) {
			cell := Cell{Alive: true}
			neighbors := 2
			
			result := cell.NextState(neighbors)
			
			if result.Alive != true {
				t.Errorf("Expected live cell with 2 neighbors to survive, got dead")
			}
		})
	})
	
	t.Run("when live cell has more than 3 neighbors", func(t *testing.T) {
		t.Run("it should die from overpopulation", func(t *testing.T) {
			cell := Cell{Alive: true}
			neighbors := 4
			
			result := cell.NextState(neighbors)
			
			if result.Alive != false {
				t.Errorf("Expected live cell with 4 neighbors to die from overpopulation, got alive")
			}
		})
	})
	
	t.Run("when live cell has fewer than 2 neighbors", func(t *testing.T) {
		t.Run("it should die from underpopulation", func(t *testing.T) {
			cell := Cell{Alive: true}
			neighbors := 1
			
			result := cell.NextState(neighbors)
			
			if result.Alive != false {
				t.Errorf("Expected live cell with 1 neighbor to die from underpopulation, got alive")
			}
		})
	})
}

func TestGrid(t *testing.T) {
	t.Run("when counting neighbors for center cell", func(t *testing.T) {
		t.Run("and grid is empty", func(t *testing.T) {
			t.Run("it should have 0 neighbors", func(t *testing.T) {
				grid := Grid{
					cells: [][]Cell{
						{{Alive: false}, {Alive: false}, {Alive: false}},
						{{Alive: false}, {Alive: false}, {Alive: false}},
						{{Alive: false}, {Alive: false}, {Alive: false}},
					},
				}
				
				neighbors := grid.CountNeighbors(1, 1)
				
				if neighbors != 0 {
					t.Errorf("Expected 0 neighbors for center cell in empty grid, got %d", neighbors)
				}
			})
		})
		
		t.Run("and has one live neighbor above", func(t *testing.T) {
			t.Run("it should have 1 neighbor", func(t *testing.T) {
				grid := Grid{
					cells: [][]Cell{
						{{Alive: false}, {Alive: true}, {Alive: false}},
						{{Alive: false}, {Alive: false}, {Alive: false}},
						{{Alive: false}, {Alive: false}, {Alive: false}},
					},
				}
				
				neighbors := grid.CountNeighbors(1, 1)
				
				if neighbors != 1 {
					t.Errorf("Expected 1 neighbor for center cell with live cell above, got %d", neighbors)
				}
			})
			
			t.Run("and has 3 live neighbors", func(t *testing.T) {
				t.Run("it should have 3 neighbors", func(t *testing.T) {
					grid := Grid{
						cells: [][]Cell{
							{{Alive: false}, {Alive: true}, {Alive: false}},
							{{Alive: true}, {Alive: false}, {Alive: true}},
							{{Alive: false}, {Alive: false}, {Alive: false}},
						},
					}
					
					neighbors := grid.CountNeighbors(1, 1)
					
					if neighbors != 3 {
						t.Errorf("Expected 3 neighbors for center cell, got %d", neighbors)
					}
				})
			})
		})
		
	})
	
	t.Run("when evolving grid to next generation", func(t *testing.T) {
		t.Run("and single live cell has no neighbors", func(t *testing.T) {
			t.Run("it should die from underpopulation", func(t *testing.T) {
				grid := Grid{
					cells: [][]Cell{
						{{Alive: false}, {Alive: false}, {Alive: false}},
						{{Alive: false}, {Alive: true}, {Alive: false}},
						{{Alive: false}, {Alive: false}, {Alive: false}},
					},
				}
				
				nextGrid := grid.NextGeneration()
				
				if nextGrid.cells[1][1].Alive != false {
					t.Errorf("Expected center cell to die from underpopulation, got alive")
				}
			})
		})
		
		t.Run("and blinker pattern (horizontal line)", func(t *testing.T) {
			t.Run("it should become vertical line", func(t *testing.T) {
				grid := Grid{
					cells: [][]Cell{
						{{Alive: false}, {Alive: false}, {Alive: false}, {Alive: false}, {Alive: false}},
						{{Alive: false}, {Alive: false}, {Alive: false}, {Alive: false}, {Alive: false}},
						{{Alive: false}, {Alive: true}, {Alive: true}, {Alive: true}, {Alive: false}},
						{{Alive: false}, {Alive: false}, {Alive: false}, {Alive: false}, {Alive: false}},
						{{Alive: false}, {Alive: false}, {Alive: false}, {Alive: false}, {Alive: false}},
					},
				}
				
				nextGrid := grid.NextGeneration()
				
				// Check vertical pattern
				if nextGrid.cells[1][2].Alive != true {
					t.Errorf("Expected cell at (1,2) to be alive in vertical blinker")
				}
				if nextGrid.cells[2][2].Alive != true {
					t.Errorf("Expected cell at (2,2) to be alive in vertical blinker")
				}
				if nextGrid.cells[3][2].Alive != true {
					t.Errorf("Expected cell at (3,2) to be alive in vertical blinker")
				}
				// Check that horizontal cells are dead
				if nextGrid.cells[2][1].Alive != false {
					t.Errorf("Expected cell at (2,1) to be dead in vertical blinker")
				}
				if nextGrid.cells[2][3].Alive != false {
					t.Errorf("Expected cell at (2,3) to be dead in vertical blinker")
				}
			})
		})
	})
}