package main

import (
	"fmt"
	"time"
)

// Worker that will run several concurrent instances
func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "stated job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		result <- j * 2
	}
}

func main() {
	fmt.Println("Start program")
	const numJobs = 2
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start up 3 workers, initially blocked because there's no jobs yey
	for w := 1; w <= 2; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs then close channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	//Collect all the result of the work
	// This "ensure" goroutine have finished
	for i := 1; i <= numJobs; i++ {
		<-results
	}

	fmt.Println("End program")
}
