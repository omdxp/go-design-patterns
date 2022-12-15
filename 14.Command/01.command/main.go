package main

import "fmt"

var overdraftLimit = -500

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Println("Deposited", amount, ", balance is now", b.balance)
}

func (b *BankAccount) Withdraw(amount int) bool {
	if b.balance-amount >= overdraftLimit {
		b.balance -= amount
		fmt.Println("Withdrew", amount, ", balance is now", b.balance)
		return true
	}
	return false
}

type Command interface {
	Call()
}

type Action int

const (
	Deposit Action = iota
	Withdraw
)

type BankAccountCommand struct {
	account *BankAccount
	action  Action
	amount  int
	success bool
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account, action, amount, false}
}

func (c *BankAccountCommand) Call() {
	switch c.action {
	case Deposit:
		c.account.Deposit(c.amount)
		c.success = true
	case Withdraw:
		c.success = c.account.Withdraw(c.amount)
	}
}

func main() {
	ba := &BankAccount{}
	cmd := NewBankAccountCommand(ba, Deposit, 100)
	cmd.Call()
	fmt.Println(ba)

	cmd = NewBankAccountCommand(ba, Withdraw, 50)
	cmd.Call()
	fmt.Println(ba)
}

// Command is a behavioral design pattern that turns a request into a stand-alone object that contains all information about the request.
// This transformation lets you parameterize methods with different requests, delay or queue a request’s execution, and support undoable operations.
