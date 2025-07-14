package main
import (
	"calming-waters/display"
	"calming-waters/model"
	"math/rand"
	"time"
)

const (
	ScreenWidth  = 80
	ScreenHeight = 20
	NumFish      = 5
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	// Initialize fish
	fish := make([]model.Fish, NumFish)
	for i := range fish {
		fish[i] = model.Fish{
			X:      rand.Intn(ScreenWidth),
			Y:      rand.Intn(ScreenHeight) + 1,
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

func play(fish []model.Fish) {
	display.ClearScreen()
	display.DrawBox(ScreenWidth, ScreenHeight)
	for i := range fish {
		fish[i].DrawFish()
		fish[i].UpdateFish()
	}
}