package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("welcom switch")
	n, _ := rand.Int(rand.Reader, big.NewInt(6))
	dice := int(n.Int64())+1
	fmt.Println(dice)
	switch dice {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
		fallthrough
	case 4:
		fmt.Println("4")
		fallthrough
	case 5:
		fmt.Println("5")
	case 6:
		fmt.Println("6")
	default:
		fmt.Println("0")

	}

}
