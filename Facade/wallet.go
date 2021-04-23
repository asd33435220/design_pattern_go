package main

import "fmt"

type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{balance: 0}
}
func (w *wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance aded successfully")
	return
}

func (w *wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("Balance is not sufficient")
	}
	fmt.Println("Wallet balance is Sufficient")
	w.balance = w.balance - amount
	return nil
}
func (w *wallet) getMoney() int {
	return w.balance
}
