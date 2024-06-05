package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to files")
	content:= "This needs to go in the file and this is the content"

	file, err:= os.Create("./myfile.txt")
    defer file.Close()

	if err!=nil {
		panic(err)
	}
    
	length, err:= io.WriteString(file, content)

	if err!=nil{
		panic(err)
	}
	fmt.Println(length)
	readfile("./myfile.txt")	

}

func readfile(filename string){
	databyte, err:= ioutil.ReadFile(filename)
	if(err!=nil){
		panic(err)
	}

	fmt.Println("Text data inside the file is : ", string(databyte))
}