package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("welcome to maths section")
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(8)+1)
	rnum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(rnum)
}
