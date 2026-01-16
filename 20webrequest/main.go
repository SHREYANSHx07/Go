package main

import (
	"fmt"
	"io"
	"net/http"
	// "os"
)

const url = "https://meet.google.com/landing"

func main() {
	fmt.Println("welcome in req of go")
	response, err := http.Get(url)
	check(err)
	fmt.Printf("type of response: %T", response)
	content, err := io.ReadAll(response.Body)
	check(err)
	fmt.Println(string(content))
	defer response.Body.Close()

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
