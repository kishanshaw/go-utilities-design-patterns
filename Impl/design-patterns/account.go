package Impl

import "fmt"

type Account struct {
	accountName string
	accountId   int16
}

func newAccount(accountId int16, accountName string) *Account {
	return &Account{
		accountId:   accountId,
		accountName: accountName,
	}
}

func (a *Account) checkAccount(accountId int16) error {
	if a.accountId == accountId {
		fmt.Println("Account verified successfully!")
		return nil
	} else {
		fmt.Println("accountId don't match")
		return fmt.Errorf("account detail is incorrect")
	}
}
