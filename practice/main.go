package main

import (
	"fmt"
	"net/url"
)

const linky = "https://google.com:3000/learn/hindi?lang=end&age=21"

func main() {
	reader,_:=url.Parse(linky)
	fmt.Println(reader.Scheme)

	linknew := &url.URL{
		Scheme: "https",
		Host: "google.com",
        Path:"/learn",
		RawQuery: "lang=hindi",
	}

	fmt.Println(linknew)

}
