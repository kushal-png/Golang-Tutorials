package main

import "fmt"

func main() {
	fmt.Println("Heyy, welcome to the tutorial on methods")
    myProfile:= User{"Kushal", 20113084, 22, false}
	fmt.Println(myProfile.GetStatus())
	fmt.Println(myProfile.NewMail())
	fmt.Println(myProfile.Name)
}

type User struct {
	Name string
	Enrollment_number int
	Age int
	Status bool
}

func (u User) GetStatus() bool {
   return u.Status
}

func (u User) NewMail() string{
     u.Name="komal"
	 return u.Name
}

// when function go inisde the classes/structs, we termed them as methods