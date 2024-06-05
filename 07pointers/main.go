package main

import "fmt"

func main() {
	fmt.Println("WELCOME TO POINTERS")
	
	var myPointer *string
	fmt.Println((myPointer))

	num:= 26
	var myPointer2 = &num

	fmt.Println(myPointer2)
	fmt.Println(*myPointer2)

}