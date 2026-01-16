package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("heelo in conversion")
	r := bufio.NewReader(os.Stdin)
	input, _ := r.ReadString('\n')
	fmt.Println("the rating is :", input)
	// newval, _ := r.ReadString('\n')
	// fmt.Println("Return yhe output", newval)
	rating, err := strconv.ParseFloat(strings.TrimSpace(input),64)
	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Println("new rating is : ", rating+1)
	}

}
