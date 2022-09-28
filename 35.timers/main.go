package main

import (
	"fmt"
	"time"
)

func main() {
	//This timer will wait 2 seconds
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C //BLOCKS on the timer's channel C until it sends a value
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer((time.Second))

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := timer2.Stop() // Can stop timer before it actually fired
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second) //This won't be trigger becase it stopped

}
