package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("welcome to race condition")
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	var score = []int{}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()
		fmt.Println("One")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Two")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Three")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
	}(wg, mut)

	wg.Wait()

	fmt.Println(score)
}
