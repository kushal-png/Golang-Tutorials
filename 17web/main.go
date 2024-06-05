package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "https://google.com"

func main() {
	fmt.Println("Welcome to web handling")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Printf("Type of response %T \n", response)
	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content:= string(databyte)
	fmt.Println(content)

}
