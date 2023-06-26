package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Concurrency using Goroutines and Channels
	fmt.Println("Concurrency example:")

	// Create a channel to receive results
	results := make(chan string)

	// Launch multiple Goroutines
	for i := 1; i <= 5; i++ {
		go processRequest(i, results)
	}

	// Wait for Goroutines to finish and collect results
	for i := 1; i <= 5; i++ {
		result := <-results
		fmt.Println(result)
	}

	// Parallelism using Goroutines and WaitGroup
	fmt.Println("\nParallelism example:")

	// Create a WaitGroup to wait for Goroutines to finish
	var wg sync.WaitGroup

	// Launch multiple Goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go processTask(i, &wg)
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	fmt.Println("\nAll Goroutines finished.")
}

func processRequest(id int, results chan<- string) {
	// Simulate some processing time
	time.Sleep(1 * time.Second)

	// Send result to the channel
	results <- fmt.Sprintf("Request %d processed", id)
}

func processTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate some processing time
	time.Sleep(1 * time.Second)

	fmt.Printf("Task %d completed\n", id)
}
