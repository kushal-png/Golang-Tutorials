package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals []string
var wg sync.WaitGroup
var mut sync.Mutex

func main() {

	websitelist := []string{
		"https://google.com",
		"https://github.com",
		"https://facebook.com",
		"https://google.com",
	}

	for _, web := range websitelist {
		go getResponse(web)
		wg.Add(1)
		// writing simply go here will not print anyhting we we didn't wait for thr threads to come back, we have to wait for them by writing waitgroup
	}

	wg.Wait()
	fmt.Println(signals)

}

func getResponse(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS")
		return
	}
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint)
}

// if you not write mutex, several goroutines are running, you don't know which goroutinw is using the memory first and in what order, mutex is mutual exclusion, it makes sure that at a single time, only one goroutine can write into the particular memory, not the other one
