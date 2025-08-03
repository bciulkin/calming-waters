package model

import (
	"calming-waters/display"
	"math/rand"
	"fmt"
)

type directionEnum int

const (
	Right directionEnum = iota
	Left
)

func direction() directionEnum {
	if (rand.Intn(2) == 1) {
		return Right
	} else {
		return Left
	}
}

type Fish struct {
	x, y   int
	template [2]string
	symbol string
	speed  int
	direction directionEnum
}

var symbols1 = [2]string {"><>", "<><"}
var symbols2 = [2]string {"><((º>", "<º))><"}


func NewFish(width, height int) Fish {
	f := Fish{
		x:         rand.Intn(width),
		y:         rand.Intn(height),
		speed:     rand.Intn(2) + 1,
		direction: direction(),
	}
	f.initTemplate()
	f.swapDirection()

	return f
}

func (f *Fish) initTemplate() {
	if (rand.Intn(2) == 0) {
		f.template = symbols1
	} else {
		f.template = symbols2
	}	
}

func (f *Fish) swapDirection() {
	if (f.direction == Right) {
		f.symbol = f.template[0]
	} else {
		f.symbol = f.template[1]
	}
}

func (f *Fish) DrawFish(width, height int) {
	if f.x < width {
		display.MoveCursor(f.y, f.x)
		fmt.Print(f.symbol)
	}
}

func (f *Fish) UpdateFish(width, height int) {
	if (f.direction == Right) {
		f.x += f.speed
	} else {
		f.x -= f.speed
	}

	// Random vertical movement
	if rand.Intn(10) > 7 {
		f.y += rand.Intn(3) - 1 // -1, 0 or +1
		if f.y < 1 {
			f.y = 1
		}
		if f.y > height {
			f.y = height
		}
	}

	// Wrap around
	if (f.x > width - 1 - len(f.symbol) && f.direction == Right) {
		f.direction = Left
		f.swapDirection()
	}

	if f.x < 1 + len(f.symbol) && f.direction == Left {
		f.direction = Right
		f.swapDirection()
	}
}
