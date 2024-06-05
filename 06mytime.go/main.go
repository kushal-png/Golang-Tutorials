package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)


	// These dates are fixed as per documentation
	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

}
