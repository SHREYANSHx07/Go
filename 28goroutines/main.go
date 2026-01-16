package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)
var signals = []string{"hello"}

var wg sync.WaitGroup
var mut sync.Mutex
func main() {
	fmt.Println("welcome in goroutines")
	website := []string{
		"https://google.com",
		"https://facebook.com",
		"https://instagram.com",
		"https://github.com",
	}
	for _,i := range website{
		go getstatuscode(i)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
	// go greet("hw")
	// greet("hello")
}

func getstatuscode(endpoint string){
	defer wg.Done()
	res,err := http.Get(endpoint)
	if err!=nil {
		log.Fatal(err)
	}
	signals = append(signals,endpoint )
	mut.Lock()
	fmt.Println(res.StatusCode,endpoint)
	mut.Unlock()
}

//  func greet(s string){
// 	for i:=0;i<6;i++{
// 		time.Sleep(3*time.Millisecond)
// 		fmt.Println(s)

// 	}
//  }
