package main

import (
	"fmt"
	"net/http"
	"time"
)


func checkStatus(urls <-chan string, resultChan chan<- string) {
    
	// Intialize the http client with one second timeout
	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	// Iterating through url channel to fetch the status
	for url := range urls {
		resp, err := client.Get(url)
		if err != nil {
			// Send error message to result channel
			resultChan <- fmt.Sprintf("%v is down", url)
		} else {
			// Close the response body to prevent resource leaks
			resp.Body.Close()
			// Send success message to result channel
			resultChan <- fmt.Sprintf("%v is up ", url)
		}

	}

}

func main() {

	const worker = 10
	urlsChan := make(chan string, len(Links))
	resultChan := make(chan string, len(Links))
	currentTime := time.Now()

	// Start workers
	for i := 0; i < worker; i++ {
		go checkStatus(urlsChan, resultChan)
	}

	// Send URLs to the urlsChan
	for _, link := range Links {
		urlsChan <- link
	}

	// Closing urls channel after sending all the urls
	close(urlsChan)

	// Collect results
	for i := 0; i < len(Links); i++ {
		fmt.Println(<-resultChan)
	}

	duration := time.Since(currentTime)
	// Print total execution time
	fmt.Printf("[ Total Time taken by Concurrent Program: %s ]", duration)
}
