package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome in post requests")
	mypostdata("http://localhost:8000/post")

}

func mypostdata(myurl string) {

	reqbody := strings.NewReader(`
	   {
	     "name" : "shreyansh",
		 "Age" : 22,
		 "mood" : "Happy"
        }
	`)
	content, err := http.Post(myurl, "application/json", reqbody)
	check(err)
	defer content.Body.Close()
	result, err := io.ReadAll(content.Body)
	check(err)
	fmt.Println(string(result))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
