package main

import "fmt"

func main(){
	fmt.Println("wlcome in functions")
	great()
	ans:=adder(1232,211221)
	fmt.Println(ans)
	ans1,str:=padder(2,32,23233,4523,2,35,55)
	fmt.Println(ans1,str)

}

func great(){
	fmt.Println("good till now")
}
func adder(x int, y int) int{
	return x+y
}
func padder(x ...int) (int,string){
	result:=0
	for i:=range x{
		result+=x[i]
	}
	return result,"goodday"
}