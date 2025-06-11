package main

import (
	"flag"
	"fmt"
	"time"

	"game-of-life-kata/pkg/gameoflife"
)

func main() {
	width := flag.Int("width", 20, "Grid width")
	height := flag.Int("height", 10, "Grid height")
	density := flag.Float64("density", 0.3, "Initial random density (0.0-1.0)")
	generations := flag.Int("generations", 50, "Number of generations to run")
	delay := flag.Int("delay", 200, "Delay between generations (milliseconds)")
	flag.Parse()

	fmt.Printf("Conway's Game of Life - %dx%d grid\n", *width, *height)
	fmt.Printf("Initial density: %.1f%%, Running %d generations\n\n", *density*100, *generations)

	game := gameoflife.NewGame(*width, *height)
	game.RandomizeGrid(*density)

	// Display initial state
	game.Display()

	// Run simulation
	for i := 0; i < *generations; i++ {
		time.Sleep(time.Duration(*delay) * time.Millisecond)
		game.Step()
		
		// Clear screen (ANSI escape sequence)
		fmt.Print("\033[2J\033[H")
		
		game.Display()
	}

	fmt.Printf("Simulation complete after %d generations.\n", *generations)
}