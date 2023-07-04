package Impl

import "fmt"

type Notification struct {
}

func (n *Notification) sendNotificationOnMoneyCredit() string {
	return fmt.Sprint("Sending Notification as money is credited")
}

func (n *Notification) sendNotificationOnMoneyDebit() string {
	return fmt.Sprint("Sending Notification as money is debited")
}
