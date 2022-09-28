package main

import (
	"fmt"
	"time"
)

func main() {
	// * Ticker for do sometghin repeatedly at intervals
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Ticker done")
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true // Tickers has been stopped so it won't receive any more values on its channel
	fmt.Println("Ticker stopped")
}
