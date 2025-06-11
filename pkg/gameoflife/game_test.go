package gameoflife

import "testing"

func TestGame(t *testing.T) {
	t.Run("when creating a new game", func(t *testing.T) {
		t.Run("it should start at generation 0", func(t *testing.T) {
			game := NewGame(5, 5)
			if game.GetGeneration() != 0 {
				t.Errorf("Expected generation 0, got %d", game.GetGeneration())
			}
		})
		
		t.Run("it should have empty grid", func(t *testing.T) {
			game := NewGame(3, 3)
			grid := game.GetGrid()
			for y := 0; y < 3; y++ {
				for x := 0; x < 3; x++ {
					if grid.GetCell(x, y) {
						t.Errorf("Expected cell (%d,%d) to be dead in new game", x, y)
					}
				}
			}
		})
	})
	
	t.Run("when stepping through generations", func(t *testing.T) {
		t.Run("it should increment generation counter", func(t *testing.T) {
			game := NewGame(5, 5)
			game.Step()
			if game.GetGeneration() != 1 {
				t.Errorf("Expected generation 1 after step, got %d", game.GetGeneration())
			}
			
			game.Step()
			if game.GetGeneration() != 2 {
				t.Errorf("Expected generation 2 after second step, got %d", game.GetGeneration())
			}
		})
	})
	
	t.Run("when randomizing grid", func(t *testing.T) {
		t.Run("and density is 0", func(t *testing.T) {
			t.Run("it should have no live cells", func(t *testing.T) {
				game := NewGame(5, 5)
				game.RandomizeGrid(0.0)
				
				grid := game.GetGrid()
				for y := 0; y < 5; y++ {
					for x := 0; x < 5; x++ {
						if grid.GetCell(x, y) {
							t.Errorf("Expected no live cells with 0 density, but found one at (%d,%d)", x, y)
						}
					}
				}
			})
		})
		
		t.Run("and density is 1", func(t *testing.T) {
			t.Run("it should have all live cells", func(t *testing.T) {
				game := NewGame(3, 3)
				game.RandomizeGrid(1.0)
				
				grid := game.GetGrid()
				for y := 0; y < 3; y++ {
					for x := 0; x < 3; x++ {
						if !grid.GetCell(x, y) {
							t.Errorf("Expected all cells to be alive with density 1.0, but found dead cell at (%d,%d)", x, y)
						}
					}
				}
			})
		})
	})
}