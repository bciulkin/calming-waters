package main
import (
	"calming-waters/display"
	"calming-waters/constants"
	"math/rand"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

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

	go func() {
		for {
			select {
			case <- done:
				return
			case <-ticker.C:
				play(fish)
			}
		}
	}()

	time.Sleep(5000 * time.Millisecond)
	ticker.Stop()
	done <- true
}

func play(fish []display.Fish) {
	display.ClearScreen()
	display.DrawBox(constants.ScreenWidth, constants.ScreenHeight)
	display.DrawFish(fish)
	display.UpdateFish(fish)
}
