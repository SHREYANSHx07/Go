package main

import (
	"fmt"
	"sync"
)

func main(){
	fmt.Println("welcome in channel section")

	wg := &sync.WaitGroup{}
	mych := make(chan int,2)
	wg.Add(2)
	go func(ch <-chan int,wg *sync.WaitGroup){
		val,isopen := <-mych
		fmt.Println(<-mych)
		fmt.Println(val)
		fmt.Println(isopen)
		//fmt.Println(<-mych)
		wg.Done()
	}(mych,wg)
	go func(ch chan<- int,wg *sync.WaitGroup){
		mych<-5
		//mych<-6
		close(mych)
		wg.Done()
	}(mych,wg)
	wg.Wait()
}
