package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome in time segment")

	presenttime := time.Now()
	fmt.Println(presenttime)
	fmt.Println(presenttime.Format("01-02-2006 15:04:05 Monday"))
	createDate := time.Date(2021, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
	//fmt.Println(presenttime.Format("01-02-2025"))

}
