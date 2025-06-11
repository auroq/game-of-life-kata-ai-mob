package gameoflife

import "testing"

func TestCell(t *testing.T) {
	t.Run("when creating a new cell", func(t *testing.T) {
		t.Run("it should be dead by default", func(t *testing.T) {
			cell := NewCell()
			if cell.IsAlive() {
				t.Error("Expected new cell to be dead, but it was alive")
			}
		})
		
		t.Run("and setting it alive", func(t *testing.T) {
			t.Run("it should be alive", func(t *testing.T) {
				cell := NewCell()
				cell.SetAlive(true)
				if !cell.IsAlive() {
					t.Error("Expected cell to be alive after setting alive to true")
				}
			})
		})
	})
}