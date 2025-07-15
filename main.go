package main

import (
	"calming-waters/display"
	"calming-waters/model"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ticker := time.NewTicker(300 * time.Millisecond)
	done := make(chan bool)

	// Initialize fish
	fish := make([]model.Fish, numFish)
	for i := range fish {
		fish[i] = model.NewFish(screenWidth, screenHeight)
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

	time.Sleep(50000 * time.Millisecond)
	ticker.Stop()
	done <- true
}

func play(fish []model.Fish) {
	display.ClearScreen()
	display.DrawBox(screenWidth, screenHeight)
	for i := range fish {
		fish[i].DrawFish()
		fish[i].UpdateFish()
	}
}
