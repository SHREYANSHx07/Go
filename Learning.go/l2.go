package main

import "fmt"

type money interface {
	balance(amount int)
}
type Creditcard struct{}
type Debittcard struct{}
type Savings struct{}

func (c Creditcard) balance(amount int) {
	fmt.Println("current balance in creditcard is :", amount)
}
func (d Debittcard) balance(amount int) {
	fmt.Println("current balance in debitcard is :", amount)
}
func (s Savings) balance(amount int) {
	fmt.Println("current balance in savings is :", amount)
}

func process(m money, amount int) {
	m.balance(amount)
}

	func main()() {
		process(Creditcard{}, 1000)
		process(Debittcard{}, 2000)
		process(Savings{}, 1000)
	}
