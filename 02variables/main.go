package main

import "fmt"

func main() {

	var username string = "kushalbnl2002"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T", username)

	var isVerified bool = true
	fmt.Println(isVerified)
	fmt.Printf("Variable is of type : %T \n", isVerified)


    // Default Values
	var num int
	fmt.Println(num)

	var num2 bool
	fmt.Println(num2)


    // no var style (Walrus Operator)
     username2:= "kushal"
	 isVerified2:= true
	 enrollment_number:= 20113084

	 fmt.Printf("%T \n", username2)
	 fmt.Printf("%T \n", isVerified2)
	 fmt.Printf("%T \n", enrollment_number)

    




}