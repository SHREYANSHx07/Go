package main

import "fmt"


type Account struct{
	balance int
}
func NewAccount(balance int) *Account {
	return &Account{balance: balance}
}

func (a *Account) deposit(money int){
	a.balance+=money
}
func (a *Account) withdraw(money int){
	a.balance-=money
}
func (a *Account) amount(){
	fmt.Println(a.balance)
}

func main(){
	a := NewAccount(555)
	a.deposit(500)
	a.withdraw(700)
	a.amount()

}

