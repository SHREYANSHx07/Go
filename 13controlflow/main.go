package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	// "strings"
)

func main() {
	fmt.Println("Welcome in if else")
	r := bufio.NewReader(os.Stdin)
	// count, _ := r.ReadString('\n')
	// count = strings.TrimSpace(count)
	// n, _ := strconv.Atoi(count)
	var n int
	fmt.Fscan(r,&n)

	var ans string
	if n <= 10 {
		ans = "low"
	} else {
		ans = "good"
	}
	fmt.Println(ans)
}
