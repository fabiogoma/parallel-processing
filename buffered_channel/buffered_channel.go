package main

import "log"

func main() {
	channel := make(chan string, 4)

	channel <- "Fabio"
	channel <- "Viviane"
	channel <- "Nicole"
	channel <- "Rafael"

	close(channel)

	for message := range channel {
		log.Println(message)
	}

}
