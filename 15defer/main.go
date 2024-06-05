package main

import "fmt"

func main() {
    defer fmt.Println("I am above both of you, AAyein main to tumse bhi niche aa gya")
	defer fmt.Println("I am above you!!, Aayein, niche kaise aa gya")
	fmt.Println("Heyy, welcome to the tutorial on defer")
}




// A defer statement invokes a function whose executiom is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or brcause the corresponding goriuotine is panicking

// works in LIFO order