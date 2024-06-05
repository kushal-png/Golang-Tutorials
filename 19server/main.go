package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("welcome to web verb video")
	PerformGetRequest()
	PerformPostRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		fmt.Println("ERROR")
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println(response.StatusCode)
	fmt.Println(response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

type Payload struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func PerformPostRequest() {
	const myurl = "http://localhost:8000/post"

	data := Payload{
		Name:  "John Doe",
		Age:   30,
		Email: "john.doe@example.com",
	}

	// Marshal the struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	//The body of the request is the JSON data contained in jsonData, which is converted to a byte buffer using bytes.NewBuffer
	response, err := http.Post(myurl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"
	data := url.Values{}
	data.Add("name", "kushal")
	data.Add("age", "22")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer response.Body.Close()
    
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
