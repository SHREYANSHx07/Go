package main

import "fmt"

func main() {
	fmt.Println("welcome in maps")

	ma := make(map[string]string)
	ma["dw"]="ercr"
	ma["crc"] = "crrcr"
	ma["re"]= "rcrce"
	fmt.Println(ma)
	fmt.Println(ma["re"])
	//delete(ma,"crc")
	fmt.Println(ma)
	for _,y := range ma {
		fmt.Println(y)
	}

}
