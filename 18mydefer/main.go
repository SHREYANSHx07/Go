package main

import "fmt"

func main(){
	fmt.Println("welcome int defer")
	defer fmt.Println("123")
	defer fmt.Println("1223")
	fun()
}

func fun(){
	for i:=0;i<=5;i++{
	    defer fmt.Println(i)
	}
}
