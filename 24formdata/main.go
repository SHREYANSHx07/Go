package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("welcome to form data urls")
	getformdata("http://localhost:8000/postform")

}

func getformdata(myurl string) {
	data := url.Values{}
	data.Add("name", "shreyansh")
	data.Add("lastname", "gupta")
	data.Add("Age", "16")
	data.Add("state", "UP")

	response, err := http.PostForm(myurl, data)
	check(err)
	fmt.Println(response)
	defer response.Body.Close()
	content , err := io.ReadAll(response.Body)
	check(err)
	fmt.Println(content)
	fmt.Println(string(content))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
