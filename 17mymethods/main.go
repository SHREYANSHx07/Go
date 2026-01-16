package main

import "fmt"

func main() {
	fmt.Println("welcome to struct")
	shreyansh := user{"Shreyansh", 22, "shreyansh1418@gmail.com", true}
	fmt.Println(shreyansh)
	fmt.Printf("the details of the usre are %+v\n",shreyansh)
	fmt.Println(shreyansh.name,shreyansh.age)
	shreyansh.getstatus()
	shreyansh.emailcon()
	fmt.Println(shreyansh)
}

type user struct {
	name  string
	age   int
	email string
	state bool
}
func (u user) getstatus(){
	fmt.Println("Is user active :",u.state)

}
func (u user) emailcon(){
	u.email="rishi@gmaail"
	fmt.Println("new mail :",u.email)
}