package main

import (
	"fmt"
	"time"
)

func main() {
	// create buffered jobs channel that will be used to receive units of work
	// and for creating a worker pool
	jobs := make(chan int, 10)

	// processed job receiver
	jobReceiver := make(chan int, 10)

	// function that creates a single worker
	createWorker := func(jobs chan int) {
		for j := range jobs {
			jobReceiver <- j
		}
	}

	// creating 10 worker pool
	for w := 1; w <= 10; w++ {
		go createWorker(jobs)
	}

	// receive some results from worker pool
	go func() {
		for {
			fmt.Println("received work unit ", <-jobReceiver)
		}
	}()

	// perform some work
	go func() {
		for {
			jobs <- 1
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			jobs <- 2
			time.Sleep(300 * time.Millisecond)
		}
	}()

	go func() {
		for {
			jobs <- 3
			time.Sleep(100 * time.Millisecond)
		}
	}()

	fmt.Scanln()
}
