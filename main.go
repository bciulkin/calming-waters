package main
import (
	"calming-waters/display"
	"calming-waters/constants"
	"math/rand"
)

func main() {
	// Initialize fish
	fish := make([]display.Fish, constants.NumFish)
	for i := range fish {
		fish[i] = display.Fish{
			X:      rand.Intn(constants.ScreenWidth),
			Y:      rand.Intn(constants.ScreenHeight) + 1,
			Symbol: []string{"><>", "><((º>"}[rand.Intn(2)],
			Speed:  rand.Intn(2) + 1,
		}
	}


	display.ClearScreen()
	display.DrawBox(constants.ScreenWidth, constants.ScreenHeight)
	display.DrawFish(fish)
	display.UpdateFish(fish)
}
