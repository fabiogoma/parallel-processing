package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}
}

func main() {
	channel := make(chan result)
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.gmcontabilidade.com.br",
		"https://www.picnic.nl",
	}

	for _, url := range urls {
		go get(url, channel)
	}

	for range urls {
		r := <-channel
		if r.err != nil {
			log.Println(r.err)
		} else {
			log.Printf("Url: %s - Time: %s\n", r.url, r.latency)
		}
	}
}
