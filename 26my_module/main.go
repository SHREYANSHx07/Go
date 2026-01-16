package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello")
	greet()
	r := mux.NewRouter()
	r.HandleFunc("/", serve).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greet() {
	fmt.Println("Hello everyone")
}

func serve(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hey there!</h1>"))

}
