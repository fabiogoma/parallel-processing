package main

import (
	"fmt"
	"sync"
)

// <- on the right side means that the channel is only used for sending data
func producer(wg *sync.WaitGroup, channel chan<- int, number int) {
	defer wg.Done()
	for idx := range number {
		fmt.Printf("PRODUCE %d\n", idx)
		channel <- idx
	}
	fmt.Println("Producer is completed!")
	close(channel)
}

// <- on the left side means that the channel is only used for receiving data
func consumer(wg *sync.WaitGroup, channel <-chan int, task string) {
	defer wg.Done()
	for value := range channel {
		fmt.Printf("consume %s %d\n", task, value)
	}
}

func main() {
	channel := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go producer(&wg, channel, 100)

	wg.Add(1)
	go consumer(&wg, channel, "task1")

	wg.Wait()

	fmt.Println("Done!")
}
