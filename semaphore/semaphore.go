package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type User struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func callApi(id int, wg *sync.WaitGroup, limit chan bool) {
	defer wg.Done()

	// Add something to the buffered channel to keep the slot occupied while the function is in progress.
	limit <- true

	// Release the buffered channel slot after the function is done, freeing up the space to another go routine
	defer func() {
		<-limit
	}()

	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", id)
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Error calling the api: %v", err)
	} else {
		defer response.Body.Close()
		var user User
		err = json.NewDecoder(response.Body).Decode(&user)
		if err != nil {
			log.Printf("Error decoding JSON for %s: %v", url, err)
		} else {
			log.Println(user.Title)
		}
	}
	// log.Println(runtime.NumGoroutine())
	// time.Sleep(time.Millisecond * 1000)
}

func main() {
	wg := new(sync.WaitGroup)

	// Create a buffered channel to limit the ammount of go routines running simultaneously
	limit := make(chan bool, 10)

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		// Spawn a new go routine, passing the limit (buffered channel) as a parameter
		go callApi(i, wg, limit)
	}

	wg.Wait()
}
