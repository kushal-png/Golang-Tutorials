package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Password string   `json:"-"`              // placing - removes that key in api
	Tags     []string `json:"tags,omitempty"` // omit empty will not include if that is empty
}

func main() {
	fmt.Println("welcome to json video")
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	ListOfCourses := []course{
		{"React", 299, "abc123", []string{"web-dev", "js"}},
		{"Angular", 29, "abc12345", []string{"full_stack", "js"}},
		{"Django", 294, "abc1234534r3",nil},
	}

	finalJson, err := json.MarshalIndent(ListOfCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", finalJson)
}

func DecodeJson() {
	jsonData := []byte(`
	{
		"coursename": "React",
		"price": 299,
		"tags": ["web-dev","js"]
    }`)

	var receivedData course
	checkValid:= json.Valid(jsonData)

	if checkValid{
		fmt.Println("Json was valid")
		json.Unmarshal(jsonData, &receivedData)
		fmt.Printf("%#v \n", receivedData)
	}


	// some cases where you want to add data to key vallue pair
    // value is of type interface because we do not sure about the type of values, which may be int or string
	var mymap map[string]interface{}
	json.Unmarshal(jsonData, &mymap)
	
	for key, value := range mymap{
		fmt.Println(key," : ", value)
	}
	
}
