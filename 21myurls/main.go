package main

import (
	"fmt"
	"net/url"
)
const myurl string = "https://csgi.wd5.myworkdayjobs.com/CSGCareers/job/India-Remote/SDE-I_31114?source=LinkedIn"

func main(){
	fmt.Println("welcome in urls")
	fmt.Println(myurl)
	u,err := url.Parse(myurl)
	check(err)
	fmt.Println(u.Scheme)
	fmt.Println(u.RawQuery)
	fmt.Println(u.Path)
	fmt.Println(u.Host)
	fmt.Println(u.Port())

	list := u.Query()
	//check(err)
	fmt.Println(list)

	for _,val := range list{
		fmt.Println(val)
	}

	partsofurl := &url.URL{
		Scheme: "https",
	    Host: "leetcode.com",
		Path: "u/shreyansh0806",
		}
	newurl := partsofurl.String()
	fmt.Println(newurl)
}


func check(err error){
	if(err!=nil){
		panic(err)
	}
}