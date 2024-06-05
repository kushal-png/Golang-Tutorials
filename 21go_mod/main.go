package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("welcome to Go mod video")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", servehome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func servehome(w http.ResponseWriter, r *http.Request) {
	// w is used to send some response back
	// r is used to read the response like urls or api's data

	w.Write([]byte("<h1>Welcome to Golang Series</h1>"))
}
