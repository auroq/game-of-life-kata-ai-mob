package gameoflife

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	grid       *Grid
	generation int
}

func NewGame(width, height int) *Game {
	return &Game{
		grid:       NewGrid(width, height),
		generation: 0,
	}
}

func (game *Game) RandomizeGrid(density float64) {
	rand.Seed(time.Now().UnixNano())
	
	for y := 0; y < game.grid.height; y++ {
		for x := 0; x < game.grid.width; x++ {
			if rand.Float64() < density {
				game.grid.SetCell(x, y, true)
			}
		}
	}
}

func (game *Game) Step() {
	game.grid = game.grid.NextGeneration()
	game.generation++
}

func (game *Game) Display() {
	fmt.Printf("Generation %d:\n", game.generation)
	fmt.Print(game.grid.String())
	fmt.Println()
}

func (game *Game) GetGeneration() int {
	return game.generation
}

func (game *Game) GetGrid() *Grid {
	return game.grid
}