package Impl

import "fmt"

type Wallet struct {
	balance int16
}

func newWallet(balance int16) *Wallet {
	return &Wallet{balance: balance}
}

func (w *Wallet) GetBalance() int16 {
	return w.balance
}

func (w *Wallet) addMoney(money int16) {
	w.balance += money
	fmt.Println("Money added to Wallet successfully")
}
func (w *Wallet) deductMoney(money int16) error {
	if w.balance >= money {
		w.balance -= money
		fmt.Println("Money debited from Wallet successfully")
		return nil
	} else {
		return fmt.Errorf("wallet's balance is not sufficient for this transaction")
	}
}
