package display

import (
	"fmt"
)

const (
	asciiRedColor   = "\033[31m"
	asciiResetColor = "\033[0m"
)

func DrawBox(width, height int) {
	if width < 2 || height < 2 {
		fmt.Println("Width and height must be at least 2")
		return
	}

	// Top border
	fmt.Print(asciiRedColor + "┌")
	for range width - 2 {
		fmt.Print("─")
	}
	fmt.Println("┐")

	// Middle rows
	for range height - 2 {
		fmt.Print("│")
		for range width - 2 {
			fmt.Print(" ")
		}
		fmt.Println("│")
	}

	// Bottom border
	fmt.Print("└")
	for range width - 2 {
		fmt.Print("─")
	}
	fmt.Println("┘" + asciiResetColor)
}

func ClearScreen() {
	fmt.Print("\033[2J") // Clear screen
	fmt.Print("\033[H")  // Move cursor to top-left
}

func MoveCursor(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
}
