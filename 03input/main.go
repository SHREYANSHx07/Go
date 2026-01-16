package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome in imputs"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the raiting")

	input, _ := reader.ReadString('\n')
	fmt.Print("thanks for reading, ", input)
	fmt.Printf("type of the raiting is %T", input)
}
