package main

import "fmt"

func main() {

	fmt.Println("Structs")
    myProfile:= User{"Kushal", 20113084, 22, false}
	fmt.Println(myProfile)
	fmt.Printf(" Details are as follows: %+v \n" ,myProfile)

}

type User struct {
	Name string
	Enrollment_number int
	Age int
	Status bool

}

// there is no inheritence in golang