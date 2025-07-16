package model

import (
	"calming-waters/display"
	"math/rand"
	"fmt"
)

type DirectionEnum int

const (
	Right DirectionEnum = iota
	Left
)

func direction() DirectionEnum {
	if (rand.Intn(2) == 1) {
		return Right
	} else {
		return Left
	}
}

type Fish struct {
	X, Y   int
	Symbol string
	Speed  int
	Direction DirectionEnum
}

func NewFish(width, height int) Fish {
	f := Fish{
		X:      rand.Intn(width - 1) + 1,
		Y:      rand.Intn(height - 1) + 1,
		Speed:  rand.Intn(2) + 1,
		Direction: direction(),
	}

	if (f.Direction == Right) {
		f.Symbol = []string{"><>", "><((º>"}[rand.Intn(2)]
	} else {
		f.Symbol = []string{"<><", "<º))><"}[rand.Intn(2)]
	}
	
	return f
}

func (f *Fish) DrawFish(width, height int) {
	if f.X < width {
		display.MoveCursor(f.Y, f.X)
		fmt.Print(f.Symbol)
	}
}

func (f *Fish) UpdateFish(width, height int) {
	if (f.Direction == Right) {
		f.X += f.Speed
	} else {
		f.X -= f.Speed
	}
	// Random vertical movement
	if rand.Intn(10) > 7 {
		f.Y += rand.Intn(3) - 1 // -1, 0 or +1
		if f.Y < 1 {
			f.Y = 1
		}
		if f.Y > height {
			f.Y = height
		}
	}
	// Wrap around
	if f.X > width {
		f.X = 0
	}
}
