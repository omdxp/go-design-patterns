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
	Undo()
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

func (c *BankAccountCommand) Undo() {
	if !c.success {
		return
	}
	switch c.action {
	case Deposit:
		c.account.Withdraw(c.amount)
	case Withdraw:
		c.account.Deposit(c.amount)
	}
}

func main() {
	ba := &BankAccount{}
	cmd := NewBankAccountCommand(ba, Deposit, 100)
	cmd.Call()
	fmt.Println(ba)

	cmd = NewBankAccountCommand(ba, Withdraw, 25)
	cmd.Call()
	fmt.Println(ba)
	cmd.Undo()
	fmt.Println(ba)
}
