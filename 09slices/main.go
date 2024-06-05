package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("WELCOME TO SLICES")
	
	var fruitlist = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("%T \n", fruitlist)
	
	fruitlist = append(fruitlist, "Mango", "Banana")
	fmt.Println(fruitlist)

    fruitlist2 := fruitlist[1:]
	fmt.Println(fruitlist)
	fmt.Println(fruitlist2)

	highScores := make([]int, 4)
	highScores[0]=234
	highScores[1]=232
	highScores[2]=246
	highScores[3]=298

	highScores= append(highScores, 666,777,321)
	fmt.Println(highScores)

	sort.Ints(highScores)
    fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	//looping
	for key, value := range(highScores){
		fmt.Println(key, " ", value)
	}
}