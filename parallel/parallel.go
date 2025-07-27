package main

import (
	"fmt"
	"sync"
)

func ParallelProcessing(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		msg := fmt.Sprintf("Go routine %s: %d", name, i)
		fmt.Println(msg)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go ParallelProcessing("Task 1", &wg)
	wg.Add(1)
	go ParallelProcessing("Task 2", &wg)
	wg.Add(1)
	go ParallelProcessing("Task 3", &wg)

	wg.Wait()
	fmt.Println("Done")
}
