package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkUrlStatus(url string, c chan string) {
	if _, err := http.Get(url); err != nil {
		// sending message to channel
		c <- url + " is down!"
	} else {
		// sending message to channel
		c <- url + " is up!"
	}
}

func main() {

	startTime := time.Now()
	// c is a channel used to commnicate from goroutines to main routine
	c := make(chan string)
	for _, url := range Links {
		go checkUrlStatus(url, c)
	}
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// Waiting for messages from channel
	for i := 0; i < len(Links); i++ {
		fmt.Println(<-c)
	}
	duration := time.Since(startTime)
	fmt.Println("Time taken by Concurrent program:", duration)
}
