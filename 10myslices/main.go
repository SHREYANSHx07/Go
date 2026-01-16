package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("wecolme to slices ")
	var a = []int{}
	a=append(a,2,3,4)
	var b = []int{}
	b=a[1:2]
	fmt.Println(a)
	fmt.Println(b)

	score := make([]int,4)
	score[0]=1
	score[1]=22222
	score[2]=23
	score[3]=11433
	fmt.Println(score)
	sort.Ints(score)
	score = append(score,233,24244)
	fmt.Println(score)


	var c = []string{"a","b","c","d","e"}
	fmt.Println(c)
	c = append(c[:2],c[3:]...)
	fmt.Println(c)
}
