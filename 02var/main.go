package main

import "fmt"
const Pans string = "shhhs"

func main() {
	// fmt.Println("Variables")
	var name string = "shreyansh"
	fmt.Println(name)
	fmt.Printf("Variable of type: %T \n",name)
	var name1 bool = true
	fmt.Println(name1)
	fmt.Printf("Variable of type: %T \n",name1)
	var name2 uint16 = 256
	fmt.Println(name2)
	fmt.Printf("Variable of type: %T \n",name2)
	var name3 float64 = 23221.34252777575766
	fmt.Println(name3)
	fmt.Printf("Variable of type: %T \n",name3)
	var name4 string
	fmt.Println(name4)
	fmt.Printf("Variable of type: %T \n",name4)

	fmt.Print(Pans)
	fmt.Printf("Variable of type: %T \n",Pans)
}
