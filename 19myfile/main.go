package main

import (
	"fmt"
	"io"
	// "io/ioutil"
	"os"
)

func main(){
	
	fmt.Println("welcome in file handling")

	s:="i am shreyansh and i am learning go file handling"
	file,err := os.Create("./filecreated")
	check(err)
	length,err:=io.WriteString(file,s)
	check(err)
	fmt.Println(length)
	defer file.Close()
	readfile("./filecreated")


}
func readfile(filename string){
	databyte, err := os.ReadFile(filename)
	check(err)
	fmt.Println(string(databyte))

}

func check(err error){
	if err!=nil{
		panic(err)
	}
}
