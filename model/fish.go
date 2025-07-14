package model

import (
	"calming-waters/display"
	"math/rand"
	"fmt"
)

type Fish struct {
	X, Y   int
	Symbol string
	Speed  int
}

func (f *Fish) DrawFish() {
	// if f.X < f.ScreenWidth {
		display.MoveCursor(f.Y, f.X)
		fmt.Print(f.Symbol)
	// }
}

func (f *Fish) UpdateFish() {
	f.X += f.Speed
	// Random vertical movement
	if rand.Intn(10) > 7 {
		f.Y += rand.Intn(3) - 1 // -1, 0 or +1
		if f.Y < 1 {
			f.Y = 1
		} 
		// else if f.Y > constants.ScreenHeight {
			// f.Y = constants.ScreenHeight
		// }
	}
	// Wrap around
	// if f.X > constants.ScreenWidth {
		// f.X = 0
	// }
}