package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {

	websitelist := []string{
		"https://google.com",
		"https://github.com",
		"https://facebook.com",
		"https://google.com",
		"https://damplips.com",
	}

	for _, web := range websitelist {
		go getResponse(web)
		wg.Add(1)
		// writing simply go here will not print anyhting we we didn't wait for thr threads to come back, we have to wait for them by writing waitgroup
	}

	wg.Wait()
}

func getResponse(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS")
		return
	}

	fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint)
}
