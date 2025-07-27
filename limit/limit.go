package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(id int, semaphore chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	// Acquire semaphore (blocks if 10 are already running)
	semaphore <- struct{}{}

	fmt.Printf("Producer %d: started\n", id)
	time.Sleep(500 * time.Millisecond) // Simulate work
	fmt.Printf("Producer %d: finished\n", id)

	// Release semaphore
	<-semaphore
}

func main() {
	const totalProducers = 100
	const maxSimultaneous = 10

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, maxSimultaneous)

	for i := range totalProducers {
		wg.Add(1)
		go producer(i, semaphore, &wg)
	}

	wg.Wait()
	fmt.Println("All producers completed!")
}
