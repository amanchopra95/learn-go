package main

import "fmt"

type BankAccount interface {
	GetBalance() int
	Deposit(amount int)
	Withdraw(amount int) error
}

func main() {
	wf := NewWellsFargo()

	wf.Deposit(100)

	currentBalance := wf.balance

	fmt.Printf("Balance %d", currentBalance)
}
