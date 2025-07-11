package display

import (
	"fmt"
	"calming-waters/constants"
	"math/rand"
)

func DrawBox(width, height int) {
	if width < 2 || height < 2 {
		fmt.Println("Width and height must be at least 2")
		return
	}

	// Top border
	fmt.Print("┌")
	for i := 0; i < width-2; i++ {
		fmt.Print("─")
	}
	fmt.Println("┐")

	// Middle rows
	for i := 0; i < height-2; i++ {
		fmt.Print("│")
		for j := 0; j < width-2; j++ {
			fmt.Print(" ")
		}
		fmt.Println("│")
	}

	// Bottom border
	fmt.Print("└")
	for i := 0; i < width-2; i++ {
		fmt.Print("─")
	}
	fmt.Println("┘")
}

func ClearScreen() {
	fmt.Print("\033[2J") // Clear screen
	fmt.Print("\033[H")  // Move cursor to top-left
}

func moveCursor(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
}

type Fish struct {
	X, Y   int
	Symbol string
	Speed  int
}

func DrawFish(fish []Fish) {
	for _, f := range fish {
		if f.X < constants.ScreenWidth {
			moveCursor(f.Y, f.X)
			fmt.Print(f.Symbol)
		}
	}
}

func UpdateFish(fish []Fish) {
	for i := range fish {
		fish[i].X += fish[i].Speed
		// Random vertical movement
		if rand.Intn(10) > 7 {
			fish[i].Y += rand.Intn(3) - 1 // -1, 0 or +1
			if fish[i].Y < 1 {
				fish[i].Y = 1
			} else if fish[i].Y > constants.ScreenHeight {
				fish[i].Y = constants.ScreenHeight
			}
		}
		// Wrap around
		if fish[i].X > constants.ScreenWidth {
			fish[i].X = 0
		}
	}
}
