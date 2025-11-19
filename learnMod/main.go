package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello world")

	// defining a router
	r := mux.NewRouter()
	r.HandleFunc("/", greetHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", r))
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World</h1>"))
}
