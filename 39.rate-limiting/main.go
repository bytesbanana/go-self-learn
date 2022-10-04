package main

import (
	"fmt"
	"time"
)

func main() {
	// Rate limiting is mechanism for controlling resource utilization and maintaining quality of servce
	// Go support ratelimiting with goroutiens, channels and tickers

	// Ex.1 Suppose we want to limit our handling of incoming requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// "limiter" channel receive a value every 200ms
	// This Rate limiting
	limiter := time.Tick(200 * time.Millisecond)

	// send 1 request every 200ms
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	//Allow short bursts of requests in rate limiting scheme while preserving the overall ratelimit
	//Can do this by "buffering" our "limiter" channel
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	//Simulate 5 more incoming request.
	//the first 3 request will benefit from the burst capability of "burstyLimiter"
	// Looking at request time in milisecond of first 3 request
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("burst request", req, time.Now())
	}

}
