package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkUrlStatus(url string) {
	/*we need to block the main goroutine until we get all responses*/
	if _, err := http.Get(url); err != nil {
		fmt.Println(url, "is down")
	} else {
		fmt.Println(url, "is up")
	}
}

func main() {
	startTime := time.Now()
	for _, url := range Links {
		checkUrlStatus(url)
	}
	duration := time.Since(startTime)
	fmt.Println("Time taken by Sequential program:", duration)

}
