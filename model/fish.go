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

const (
	symbols = [4]string {"><>", "<><", "><((º>", "<º))><"}
)

func NewFish(width, height int) Fish {
	f := Fish{
		X:         rand.Intn(width) + 1,
		Y:         rand.Intn(height) + 1,
		Speed:     rand.Intn(2) + 1,
		Direction: direction(),
	}
	f.swapDirection()

	return f
}

func (f *Fish) swapDirection() {
	if (f.Direction == Right) {
		f.Symbol = symbols[rand.Intn(2)]
	} else {
		f.Symbol = symbols[rand.Intn(2) + 2]
	}
	
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
	if (f.X > width - 1 - len(f.Symbol)) {
		f.Direction = Left
		f.swapDirection()
	}

	if f.X < 1 {
		f.Direction = Right
		f.swapDirection()
	}
}
