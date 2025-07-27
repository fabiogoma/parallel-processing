package main

import (
	"log"
)

func main() {
	channel := make(chan int, 1)
	channel <- 42

	a, ok := <-channel

	// 42 - Channel content
	// true - Channel is open
	log.Println(a, ok)

	close(channel)

	// 0 - Channel default value (int default is 0)
	// false - Channel is closed
	b, ok := <-channel
	log.Println(b, ok)
}
