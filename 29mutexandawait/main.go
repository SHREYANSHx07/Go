package main

import (
	"fmt"
	"sync"
)

func main() {
	var s = []int{0}
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	fmt.Println("welcome to use of await and mutex")
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("1st goroutine")
		mut.Lock()
		s = append(s, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("2nd goroutine")
		mut.Lock()
		s = append(s, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("3rd goroutine")
		mut.Lock()
		s = append(s, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("4th goroutine")
		mut.RLock()
		fmt.Println(s)
		mut.RUnlock()
		//wg.Done()
	}(wg, mut)
	wg.Wait()
	fmt.Println(s)
}
