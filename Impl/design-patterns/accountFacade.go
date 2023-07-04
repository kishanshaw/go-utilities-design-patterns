package Impl

import "fmt"

type AccountFacade struct {
	Account      *Account
	Wallet       *Wallet
	Notification *Notification
}

func NewAccountFacade(accountId int16, accountName string, balance int16) *AccountFacade {
	fmt.Println("account creation in progress")
	accountFacade := &AccountFacade{
		Account:      newAccount(accountId, accountName),
		Wallet:       newWallet(balance),
		Notification: &Notification{},
	}
	fmt.Println("account creation completed")
	return accountFacade
}

func (a *AccountFacade) AddMoneyToWallet(accountId int16, money int16) error {
	fmt.Println("starting to add money to account's wallet")
	err := a.Account.checkAccount(accountId)
	if err != nil {
		return err
	}
	a.Wallet.addMoney(money)
	fmt.Println(a.Notification.sendNotificationOnMoneyCredit())
	return nil
}

func (a *AccountFacade) DeductMoneyFromWallet(accountId int16, money int16) error {
	fmt.Println("Starting to deduct money from account's wallet")
	err := a.Account.checkAccount(accountId)
	if err != nil {
		return err
	}
	debitError := a.Wallet.deductMoney(money)
	if debitError != nil {
		return debitError
	}
	fmt.Println(a.Notification.sendNotificationOnMoneyDebit())
	return nil
}
