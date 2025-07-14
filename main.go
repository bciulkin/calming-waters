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



	go func() {
		for {
			select {
			case <- done:
				return
			case <-ticker.C:
				play()
			}
		}
	}()

	time.Sleep(5000 * time.Millisecond)
	ticker.Stop()
	done <- true
}

func play() {
	display.ClearScreen()
	display.DrawBox(constants.ScreenWidth, constants.ScreenHeight)
	display.DrawFish()
	display.UpdateFish()
}
