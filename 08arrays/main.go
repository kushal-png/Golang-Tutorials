package main

import "fmt"

func main() {
	fmt.Println("WELCOME TO ARRAYS")
	
	var fruitlist [4]string
	fruitlist[0]="Apple"
	fruitlist[1]="tomato"
	fruitlist[2]="peach"

	fmt.Println(fruitlist)
	fmt.Println(len(fruitlist))

	var veglist = [3]string{"potato", "tomato","mushroom"}
	fmt.Println(veglist)


	// looping
	for key, value := range(veglist){
		fmt.Println(key, " ", value)
	}

}