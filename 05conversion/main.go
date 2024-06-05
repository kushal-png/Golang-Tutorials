package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza Rating app")
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating from 1 to 5")
	rating, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating :", rating)

	// Input rating is of type string, we have t0 convert it into number
	fmt.Printf("%T \n", rating)

	// note that trailing character also got added into the string (\n), therefore we need to trim that first
	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(numRating+1)
	}
}
