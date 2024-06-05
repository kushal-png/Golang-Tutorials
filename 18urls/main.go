package main

import (
	"fmt"
	"net/url"
)

const urll string = "https://google.com:3000/learn?coursename=reactjs&paymentid=94rfc87cgt8"

func main() {
	fmt.Println("Welcome to urls")

	//Parts of URL
	result, _ := url.Parse(urll)
    fmt.Printf("%T \n",result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println(qparams)

	//Construct a url from thr parts

	partsofUrl:= &url.URL{
		Scheme: "https",
		Host: "google.com",
		Path: "/learn/me",
		RawQuery: "user=kushal",
	}

	anotherUrl:= partsofUrl.String()
	fmt.Println(anotherUrl)
}
