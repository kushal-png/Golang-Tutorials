package main

import "fmt"

func main() {
	fruitlist := make([]string, 4)
	fmt.Println(fruitlist)

	fruitlist[0]="apple"
	fruitlist[1]="mango"
	fruitlist[2]="banana"
	fruitlist[3]="orange"
    fmt.Println(fruitlist)

	//delete by index
	index:=2
	fruitlist=append(fruitlist[:index], fruitlist[index+1:]...)
    fmt.Println(fruitlist)
}
