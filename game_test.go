package main

import "testing"

func TestGameOfLife(t *testing.T) {
	t.Run("when cell is dead", func(t *testing.T) {
		t.Run("and has no neighbors", func(t *testing.T) {
			t.Run("it should remain dead", func(t *testing.T) {
				cell := false
				result := nextState(cell, 0)
				if result != false {
					t.Errorf("Expected false, got %v", result)
				}
			})
		})
		
		t.Run("and has exactly 3 neighbors", func(t *testing.T) {
			t.Run("it should become alive", func(t *testing.T) {
				cell := false
				result := nextState(cell, 3)
				if result != true {
					t.Errorf("Expected true, got %v", result)
				}
			})
		})
	})
	
	t.Run("when cell is alive", func(t *testing.T) {
		t.Run("and has 2 neighbors", func(t *testing.T) {
			t.Run("it should survive", func(t *testing.T) {
				cell := true
				result := nextState(cell, 2)
				if result != true {
					t.Errorf("Expected true, got %v", result)
				}
			})
		})
		
		t.Run("and has 3 neighbors", func(t *testing.T) {
			t.Run("it should survive", func(t *testing.T) {
				cell := true
				result := nextState(cell, 3)
				if result != true {
					t.Errorf("Expected true, got %v", result)
				}
			})
		})
		
		t.Run("and has 1 neighbor", func(t *testing.T) {
			t.Run("it should die from underpopulation", func(t *testing.T) {
				cell := true
				result := nextState(cell, 1)
				if result != false {
					t.Errorf("Expected false, got %v", result)
				}
			})
		})
		
		t.Run("and has 4 neighbors", func(t *testing.T) {
			t.Run("it should die from overpopulation", func(t *testing.T) {
				cell := true
				result := nextState(cell, 4)
				if result != false {
					t.Errorf("Expected false, got %v", result)
				}
			})
		})
	})
}

func TestNeighborCounting(t *testing.T) {
	t.Run("when counting neighbors in 3x3 grid", func(t *testing.T) {
		t.Run("and center cell has all neighbors alive", func(t *testing.T) {
			t.Run("it should count 8 neighbors", func(t *testing.T) {
				grid := [][]bool{
					{true, true, true},
					{true, false, true},
					{true, true, true},
				}
				count := countNeighbors(grid, 1, 1)
				if count != 8 {
					t.Errorf("Expected 8, got %v", count)
				}
			})
		})
		
		t.Run("and center cell has 3 neighbors alive", func(t *testing.T) {
			t.Run("it should count 3 neighbors", func(t *testing.T) {
				grid := [][]bool{
					{true, true, true},
					{false, false, false},
					{false, false, false},
				}
				count := countNeighbors(grid, 1, 1)
				if count != 3 {
					t.Errorf("Expected 3, got %v", count)
				}
			})
		})
		
		t.Run("and top-left corner has 3 possible neighbors alive", func(t *testing.T) {
			t.Run("it should count 3 neighbors", func(t *testing.T) {
				grid := [][]bool{
					{false, true, false},
					{true, true, false},
					{false, false, false},
				}
				count := countNeighbors(grid, 0, 0)
				if count != 3 {
					t.Errorf("Expected 3, got %v", count)
				}
			})
		})
	})
}

func TestGridEvolution(t *testing.T) {
	t.Run("when evolving a 3x3 grid", func(t *testing.T) {
		t.Run("and center cell should reproduce", func(t *testing.T) {
			t.Run("it should create new generation", func(t *testing.T) {
				grid := [][]bool{
					{true, true, true},
					{false, false, false},
					{false, false, false},
				}
				expected := [][]bool{
					{false, true, false},
					{false, true, false},
					{false, false, false},
				}
				result := evolveGrid(grid)
				for i := range result {
					for j := range result[i] {
						if result[i][j] != expected[i][j] {
							t.Errorf("At [%d][%d]: expected %v, got %v", i, j, expected[i][j], result[i][j])
						}
					}
				}
			})
		})
	})
}