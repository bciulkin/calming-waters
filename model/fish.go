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

var symbols1 = [2]string {"><>", "<><"}
var symbols2 = [2]string {"><((º>", "<º))><"}


func NewFish(width, height int) Fish {
	f := Fish{
		X:         rand.Intn(width),
		Y:         rand.Intn(height),
		Speed:     rand.Intn(2) + 1,
		Direction: direction(),
	}
	f.initSymbol()

	return f
}

func (f *Fish) initSymbol() {
	if (f.Direction == Right) {
		f.Symbol = symbols1[rand.Intn(1)]
	} else {
		f.Symbol = symbols1[rand.Intn(1) + 1]
	}
	
}

func (f *Fish) swapDirection() {
	if (f.Direction == Right) {
		f.Symbol = symbols1[rand.Intn(1)]
	} else {
		f.Symbol = symbols1[rand.Intn(1) + 1]
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
	if (f.X > width - 1 - len(f.Symbol) && f.Direction == Right) {
		f.Direction = Left
		f.swapDirection()
	}

	if f.X < 1 + len(f.Symbol) && f.Direction == Left {
		f.Direction = Right
		f.swapDirection()
	}
}
