package main

import "fmt"

func main() {
	fmt.Println("Heyy, welcome to the tutorial on functions")
    greet()

	result:= adder(3,5)
	fmt.Println(result)

	result2:= proadder(3,5,1,2,3,4,5)
	fmt.Println(result2)
}


// simple function
func greet(){
	fmt.Println("Hello, i am greeter")
}

//complex
func adder( a int, b int) int {
	return a+b
}

func proadder(values ...int)int{
	result:=0
	for _,value:= range(values){
       result+=value
	}
	return result
}