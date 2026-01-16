package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome in web get req")
	myurl := "http://localhost:8000"
	responsecode(myurl)

}

func responsecode(myurl string) {
	
	response, err := http.Get(myurl)
	check(err)
	fmt.Println(response.Status)
	fmt.Println(response.Request.Method)
	fmt.Println(response.ContentLength)
	defer response.Body.Close()
	// content, err := io.ReadAll(response.Body)
	// fmt.Println(string(content))

	var newstr strings.Builder
	content, err := io.ReadAll(response.Body)
	bytecontent, _ := newstr.Write(content)
	fmt.Println(bytecontent)
	fmt.Println(newstr.String())
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
