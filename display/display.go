package display

import (
	"fmt"
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

func MoveCursor(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
}