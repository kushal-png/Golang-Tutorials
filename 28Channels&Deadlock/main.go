// As of now, only waitgroup works as a manager, however goroutines do not that any other goroutine is there or not, this info is only with waitgroup.
// Channels:- Are a way in which your multiple goroutines can actually talk to each other

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("This is about CHannels")
	wg := &sync.WaitGroup{}
	ch := make(chan int,1)

	wg.Add(2)
	go func(wg *sync.WaitGroup, ch chan int) {
		fmt.Println(<-ch)
		defer wg.Done()

	}(wg, ch)
	go func(wg *sync.WaitGroup, ch chan int) {
		ch <- 5
		ch <-6
		close(ch)
		defer wg.Done()
	}(wg, ch)
	wg.Wait()
}
