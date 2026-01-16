package main

import "fmt"

func main() {
	fmt.Println("welcome to struct")
	shreyansh := user{"Shreyansh", 22, "shreyansh1418@gmail.com", true}
	fmt.Println(shreyansh)
	fmt.Printf("the details of the usre are %+v\n",shreyansh)
	fmt.Println(shreyansh.name,shreyansh.age)
}

type user struct {
	name  string
	age   int
	email string
	state bool
}
