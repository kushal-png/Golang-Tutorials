package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware, helper
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	rou := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{FullName: "kushal", Website: "helloworld"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "MERN", CoursePrice: 199, Author: &Author{FullName: "kushal_jindAL", Website: "hello__world"}})

	//routing

	rou.HandleFunc("/", servehome).Methods("GET")
	rou.HandleFunc("/courses", getAllCourses).Methods("GET")
	rou.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	rou.HandleFunc("/create", createOneCourse).Methods("POST")
	rou.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	rou.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to port
	log.Fatal(http.ListenAndServe(":4000", rou))

}

//controllers

//servehomeroute

func servehome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to API by Kushal Jindal"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "Application/json")

	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one courses")
	w.Header().Set("Content-Type", "Application/json")

	//grab id from request
	params := mux.Vars(r)

	// loop through the courses, find matching id and return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course find with given id")

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one courses")
	w.Header().Set("Content-Type", "Application/json")

	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what about -{}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}

	// generate a unique id, string & append taht in course
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one courses")
	w.Header().Set("Content-Type", "Application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var upadteCourse Course
			json.NewDecoder(r.Body).Decode(&upadteCourse)
			upadteCourse.CourseId = params["id"]

			courses = append(courses, upadteCourse)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// send a response when id is not found

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one courses")
	w.Header().Set("Content-Type", "Application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

	// send a response when id is not found

}
