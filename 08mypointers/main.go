package main

import "fmt"

func main() {
	fmt.Println("welcome in pointers")
	// var ptr *int 
	// fmt.Println(ptr)
	mynum :=23
	var ptr = &mynum;
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr * 2

	fmt.Println(mynum)
}
