package main

import "fmt"

func main(){
	fmt.Println("welcome in loops")

	days := []string{"monday","tuesday","wedsday","thrusday","friday"}
	fmt.Println(days)

	// for i:=0;i<len(days);i++{
	// 	fmt.Printf("%v ",days[i])
	// }

	// for i:=range days{
	// 	fmt.Println(days[i])
	// }

	for i,j := range days{
		fmt.Println("index is :",i,"the valuse is",j)
	}

	value := 12
	 for value<=20{
		if value==14{
			goto xyz
		}
		if value==15{
			break
		}
		fmt.Println(value)
		value++
	 }
	
	 xyz:
	 fmt.Println("learrning goto")
	

}
